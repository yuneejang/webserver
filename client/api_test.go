package client

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/yuneejang/webserver/types"
)

func TestJSONRequestParsing(t *testing.T) {
	mem := types.Member{
		Name:   "Alex",
		Age:    10,
		Active: true,
	}
	fmt.Println(mem.Name, mem.Age, mem.Active)

	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}
	fmt.Println("marshal jsonByte : ", "jsonBytes", jsonBytes)
	fmt.Println("marshal jsonByte (string) : ", "jsonBytes", (string)(jsonBytes))

	var decoded *types.Member
	err = json.Unmarshal(jsonBytes, &decoded)
	if err != nil {
		panic(err)
	}

	fmt.Println("unmarshal jsonByte : ", "decoded", decoded)

}
