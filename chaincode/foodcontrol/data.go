package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

//Descripcion del activo

type Data struct {
	Column1 string `json:"colum1"`
	Column2 string `json:"column2"`
}

func (s *SmartContract) Set(ctx contractapi.TransactionContextInterface, id string, column1 string, column2 string) error {

	// validaciones, que no exista el food id

	data := Data{
		Column1: column1,
		Column2: column2,
	}

	dataAsBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Marshal error: %s", err.Error())
		return err
	}

	return ctx.GetStub().PutState(id, dataAsBytes)
}

func (s *SmartContract) Query(ctx contractapi.TransactionContextInterface, id string) (*Data, error) {

	dataAsBytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world  state %s", err.Error())
	}

	if dataAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", id)
	}

	data := new(Data)

	err = json.Unmarshal(dataAsBytes, data)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error %s", err.Error())
	}

	return data, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create data chaincode %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error create data chaincode %s", err.Error())
	}
}
