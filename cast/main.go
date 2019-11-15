package main

import (
	"fmt"
	"reflect"

	"encoding/json"
	// "github.com/spf13/cast"
)

func main() {
	/*
		i := `1`
		s, err := cast.ToInt64E(i)
		if nil != err {
			fmt.Println(err)
			return
		}
		fmt.Println(reflect.TypeOf(s), s)
		bt, err := json.Marshal(i)
		fmt.Println(fmt.Sprintf("%s", bt), err)

		var j interface{}
		err = json.Unmarshal(bt, &j)
		fmt.Println(err)
		fmt.Println(reflect.TypeOf(i), reflect.TypeOf(j))
	*/

	bt := []byte(`{"id":"febad709-2a41-4f96-94dc-a159298e3bc9","type":"request-size-limiting","config":{"allowed_payload_size":1}}`)

	var strategy map[string]interface{}

	fmt.Println(json.Unmarshal(bt, &strategy))
	config, ok := strategy["config"].(map[string]interface{})

	fmt.Println(config, ok)
	size := config["allowed_payload_size"]
	fmt.Println(config["allowed_payload_size"], reflect.TypeOf(size))
	switch size.(type) {
	case int64:
		fmt.Println("int64")
	case float64:
		fmt.Println("float64")
	}
}
