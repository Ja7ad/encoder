# encoder
create multiple encoder object for encode and decode data and an example for register pattern.

## install

```shell
go get -u github.com/Ja7ad/encoder
```

## usage

example usage for encoder :

```go
package main

import (
	"fmt"
	"log"
	"github.com/Ja7ad/encoder"
)

type Person struct {
	Name    string `json:"name" bson:"name"`
	Age     int    `json:"age" bson:"age"`
	Address string `json:"address" bson:"address"`
}

func main() {
	enc := New()
	enc.RegisterEncoder(JSON, &JsonEncoder{})
	enc.RegisterEncoder(GOB, &GobEncoder{})
	enc.RegisterEncoder(BSON, &BsonEncoder{})
	enc.RegisterEncoder(PROTO, &ProtoEncoder{})

	p := &Person{}

	jsonEnc, err := enc.GetJsonEncoder()
	if err != nil {
		log.Fatal(err)
	}

	b, err := jsonEnc.Encode(&Person{
		Name:    "Javad",
		Age:     30,
		Address: "example address 1",
	})

	if err != nil {
		log.Fatal(err)
	}

	if err = jsonEnc.Decode(b, p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}
```
