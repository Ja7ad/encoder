package encoder

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/proto"
	"sync"
)

type EncodeType int

const (
	JSON  EncodeType = iota // JSON is encoder and decoder for json
	GOB                     // GOB is encoder and decoder for gob
	BSON                    // BSON is encoder and decoder for mongodb bson
	PROTO                   // PROTO is encoder and decoder for proto reflect
)

type Encode struct {
	encoders map[EncodeType]Encoder
	lk       sync.Mutex
}

type Encoder interface {
	Encode(data interface{}) ([]byte, error)
	Decode(data []byte, vPtr interface{}) error
}

// New create encode object
func New() *Encode {
	return &Encode{
		encoders: make(map[EncodeType]Encoder),
		lk:       sync.Mutex{},
	}
}

// RegisterEncoder register encode base on EncodeType and Encoder interface
func (e *Encode) RegisterEncoder(encoderType EncodeType, encoder Encoder) {
	e.lk.Lock()
	defer e.lk.Unlock()
	e.encoders[encoderType] = encoder
}

func (e *Encode) GetJsonEncoder() (Encoder, error) {
	if v, ok := e.encoders[JSON]; ok {
		return v, nil
	}
	return nil, errors.New("json encode not registered")
}

func (e *Encode) GetGobEncoder() (Encoder, error) {
	if v, ok := e.encoders[GOB]; ok {
		return v, nil
	}
	return nil, errors.New("gob encode not registered")
}

func (e *Encode) GetBsonEncoder() (Encoder, error) {
	if v, ok := e.encoders[BSON]; ok {
		return v, nil
	}
	return nil, errors.New("bson encode not registered")
}

func (e *Encode) GetProtoEncoder() (Encoder, error) {
	if v, ok := e.encoders[PROTO]; ok {
		return v, nil
	}
	return nil, errors.New("proto encode not registered")
}

type JsonEncoder struct{}

func (*JsonEncoder) Encode(vPtr interface{}) ([]byte, error) {
	return json.Marshal(vPtr)
}

func (*JsonEncoder) Decode(data []byte, vPtr interface{}) error {
	return json.Unmarshal(data, vPtr)
}

type GobEncoder struct{}

func (*GobEncoder) Encode(vPtr interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(vPtr); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (*GobEncoder) Decode(data []byte, vPtr interface{}) error {
	dec := gob.NewEncoder(bytes.NewBuffer(data))
	return dec.Encode(vPtr)
}

type BsonEncoder struct{}

func (*BsonEncoder) Encode(vPtr interface{}) ([]byte, error) {
	return bson.Marshal(vPtr)
}

func (*BsonEncoder) Decode(data []byte, vPtr interface{}) error {
	return bson.Unmarshal(data, vPtr)
}

type ProtoEncoder struct{}

func (*ProtoEncoder) Encode(vPtr interface{}) ([]byte, error) {
	if v, ok := vPtr.(proto.Message); ok {
		return proto.Marshal(v)
	}

	return nil, errors.New("data isn't proto message")
}

func (*ProtoEncoder) Decode(data []byte, vPtr interface{}) error {
	if _, ok := vPtr.(*interface{}); ok {
		return nil
	}

	if v, ok := vPtr.(proto.Message); ok {
		return proto.Unmarshal(data, v)
	}

	return errors.New("vPtr is not proto message")
}
