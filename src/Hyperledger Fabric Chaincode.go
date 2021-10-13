package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"strconv"
)

type CC struct {
}

func (c *CC) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// TODO: initialize states of a, b, c, d, bank
	err := stub.PutState("a", []byte("100"))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState("b", []byte("100"))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState("c", []byte("100"))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState("d", []byte("100"))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState("bank", []byte("1000"))
	if err != nil {
		return shim.Error(err.Error())
	}
	
	return shim.Success([]byte("OK"))
}

func (c *CC) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	var f, args = stub.GetFunctionAndParameters()
	switch f {
	case "init":
		return c.Init(stub)
	case "send":

		money, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			return shim.Error(err.Error())
		}
		aValByte, err := stub.GetState(args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		bValByte, err := stub.GetState(args[1])
		if err != nil {
			return shim.Error(err.Error())
		}

		aVal, err := strconv.ParseFloat(string(aValByte), 64)
		if err != nil {
			return shim.Error(err.Error())
		}

		bVal, err := strconv.ParseFloat(string(bValByte), 64)
		if err != nil {
			return shim.Error(err.Error())
		}

		aVal -= money
		bVal += money

		if aVal <0 {
			return shim.Error(errors.New("NON-POSITIVE BALANCE").Error())
		}

		aValUpdatedByte := fmt.Sprintf("%f", aVal)
		bValUpdatedByte := fmt.Sprintf("%f", bVal)

		if err = stub.PutState(args[0], []byte(aValUpdatedByte)); err != nil {
			return shim.Error(err.Error())
		}
		if err = stub.PutState(args[1], []byte(bValUpdatedByte)); err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success([]byte(args[0] + ":" + aValUpdatedByte + ", " + args[1] + ":" + bValUpdatedByte))
	case "withdraw":
		// TODO: Implement a withdraw function, e.g., {withdraw a 20 0.05}.
		money, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			return shim.Error(err.Error())
		}
		fee, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			return shim.Error(err.Error())
		}
		aValByte, err := stub.GetState(args[0])
		if err != nil {
			return shim.Error(err.Error())
		}
		bValByte, err := stub.GetState("bank")
		if err != nil {
			return shim.Error(err.Error())
		}

		aVal, err := strconv.ParseFloat(string(aValByte), 64)
		if err != nil {
			return shim.Error(err.Error())
		}

		bVal, err := strconv.ParseFloat(string(bValByte), 64)
		if err != nil {
			return shim.Error(err.Error())
		}

		aVal -= money * (1 + fee)
		bVal += money * fee

		if aVal <0 {
			return shim.Error(errors.New("NON-POSITIVE BALANCE").Error())
		}

		aValUpdatedByte := fmt.Sprintf("%f", aVal)
		bValUpdatedByte := fmt.Sprintf("%f", bVal)

		if err = stub.PutState(args[0], []byte(aValUpdatedByte)); err != nil {
			return shim.Error(err.Error())
		}
		if err = stub.PutState("bank", []byte(bValUpdatedByte)); err != nil {
			return shim.Error(err.Error())
		}
		// TODO: For the below return statement, you should replace ?? with your variable names.
		return shim.Success([]byte(args[0] + ":" + aValUpdatedByte + ", " + "bank" + ":" + bValUpdatedByte))
	}
	return shim.Error("No function is supported for " + f)
}

func main() {
	err := shim.Start(new(CC))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Start simple chaincode now")
}
