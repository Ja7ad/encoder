package encoder

import (
	"github.com/Ja7ad/encoder/testdata"
	"testing"
)

type Person struct {
	Name    string `json:"name" bson:"name"`
	Age     int    `json:"age" bson:"age"`
	Address string `json:"address" bson:"address"`
}

func Test_Encoder(t *testing.T) {
	t.Parallel()

	enc := New()
	enc.RegisterEncoder(JSON, &JsonEncoder{})
	enc.RegisterEncoder(GOB, &GobEncoder{})
	enc.RegisterEncoder(BSON, &BsonEncoder{})
	enc.RegisterEncoder(PROTO, &ProtoEncoder{})

	t.Run("json encoder", func(t *testing.T) {
		p := &Person{}

		jsonEnc, err := enc.GetJsonEncoder()
		if err != nil {
			t.Fatal(err)
		}

		b, err := jsonEnc.Encode(&Person{
			Name:    "Javad",
			Age:     30,
			Address: "example address 1",
		})

		if err != nil {
			t.Fatal(err)
		}

		if err = jsonEnc.Decode(b, p); err != nil {
			t.Fatal(err)
		}

		t.Log(p)
	})

	t.Run("gob encoder", func(t *testing.T) {
		p := &Person{}

		gobEnc, err := enc.GetGobEncoder()
		if err != nil {
			t.Fatal(err)
		}

		b, err := gobEnc.Encode(&Person{
			Name:    "Ali",
			Age:     30,
			Address: "example address 2",
		})

		if err != nil {
			t.Fatal(err)
		}

		if err = gobEnc.Decode(b, p); err != nil {
			t.Fatal(err)
		}

		t.Log(p)
	})

	t.Run("bson encoder", func(t *testing.T) {
		p := &Person{}

		bsonEnc, err := enc.GetBsonEncoder()
		if err != nil {
			t.Fatal(err)
		}

		b, err := bsonEnc.Encode(&Person{
			Name:    "Saeed",
			Age:     30,
			Address: "example address 3",
		})

		if err != nil {
			t.Fatal(err)
		}

		if err = bsonEnc.Decode(b, p); err != nil {
			t.Fatal(err)
		}

		t.Log(p)
	})

	t.Run("proto encoder", func(t *testing.T) {
		p := &testdata.Person{}

		protoEnc, err := enc.GetProtoEncoder()
		if err != nil {
			t.Fatal(err)
		}

		b, err := protoEnc.Encode(&testdata.Person{
			Name:    "Ali",
			Age:     30,
			Address: "example address 2",
		})

		if err != nil {
			t.Fatal(err)
		}

		if err = protoEnc.Decode(b, p); err != nil {
			t.Fatal(err)
		}

		t.Log(p)
	})

}
