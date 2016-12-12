package main

import (
	"errors"
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// DisputeClaim request structure
type DisputeClaimRecord struct{

	DisputeId string `json:"claimid"`
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
	
	fmt.Println("Init is running ")
	
	return nil, nil
}

// Invoke is an entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	fmt.Println("invoke is running " + function)

	if function == "write" {
		var key string
		if len(args) != 5 {
			return nil, errors.New("Incorrect number of arguments. Expecting 5. name of the key and value to set")
		}
		key = args[0]
		
		disputeRecord := &DisputeClaimRecord{ DisputeId: args[1], TransactionId: args[2], DisputeType: args[3], Comments: args[4]}
		disputeRecordJSON, err := json.Marshal(disputeRecord)
		
		if(err != nil){
			fmt.Println("Error while creating JSON structure: %s" , err)		
		}
		
		// store the JSON on ledger
		err = stub.PutState(key, disputeRecordJSON) //write the variable into the chaincode state
		if err != nil {
			return nil, err
		}		
	}
		return nil, nil
}

// Query is an entry point for queries
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