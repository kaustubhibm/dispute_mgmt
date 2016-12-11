package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// DisputeClaim request structure
type DisputeClaim struct{

	ClaimId string `json:"claimid"`
	TransactionId string `json:"transactionid"`
	DisputeType string `json:"disputetype"`
	Comments string `json:"comments"`
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Dispute Management chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	if function == "write" {
		var key, value string
		var err error
		if len(args) != 2 {
			return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
		}
		key = args[0]
		value = "{\"" + args[1] + "\"," + args[2] + "\"," + args[3] + "\"," + args[4] + "\""
		
		fmt.Println("********>>>>>> Value is " + value)
		
		err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
		if err != nil {
			return nil, err
		}
	}
		return nil, nil
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
	
	if function == "read" {
	
		var key, jsonResp string
		var err error

		if len(args) != 1 {
			return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
		}

		key = args[0]
		valAsbytes, err := stub.GetState(key)
		if err != nil {
			jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
			return nil, errors.New(jsonResp)
		}
		
		return valAsbytes, nil
	}
	return nil, nil
}