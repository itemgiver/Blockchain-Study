package main

import (
	"./kv"
	"fmt"
	"os"
)

func main() {
	db := &kv.Database{}
	kv.Init(db)
	defer kv.Finalize(db)

	for i, arg := range os.Args[1:] {
		fmt.Println(i, "TestPutState("+arg+"_key,"+arg+"_value): ", db.PutState(arg+"_key", arg+"_value"))
	}

	for i, arg := range os.Args[1:] {
		value, err := db.GetState(arg + "_key")
		fmt.Println(i, "TestGetState("+arg+"_key): ", value, err)
	}
}
