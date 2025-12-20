package pkg

import (
	"GoChat/internal/logger"
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
)

/*
Here we Define the Transmission Struct, This Struct is used to encapsulate all data sent between Clients, Servers, and ClientHandlers
This provides a standard type that we can decode all bytecode transmissions into before seeing what type the data actually is VIA the DataCode

We can just encode the entire transmission as long as the inner data type is registered in the gob.go init function. 
each time we decode a transmission we will have to read the DataCode and then use that information to decide how to interpret the inner data*/
type Transmission struct {
	Code DataCode
	Data Data  
}

func CreateTransmission(DC DataCode, D Data) *Transmission {
	return &Transmission{Code: DC, Data: D}
}


func (T *Transmission) EncodeTransmission() ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)



	err := encoder.Encode(T)
	if err != nil {
		return nil, err
	}
	logger.Logger.Info("Encoded Transmission", "Data", reflect.TypeOf(T.Data), "Num Bytes", len(buffer.Bytes()))

	return buffer.Bytes(), nil
}


func DecodeTransmission(data []byte) (*Transmission, error) {
	var buffer bytes.Buffer
	var DecodedTransmission Transmission
	buffer.Write(data)
	decoder := gob.NewDecoder(&buffer)
	err := decoder.Decode(&DecodedTransmission) 
	if err != nil {
		fmt.Println("FAILED IN DECODE")	
		return nil, err
	}

	logger.Logger.Info("Decoded Transmission", "DataCode", DecodedTransmission.Code)
	return &DecodedTransmission, nil	
}

//--------------------------------------------------------------------
//--------------------------------------------------------------------	
type DataCode uint8 
const ( TestDataCode DataCode = iota) 
type Data interface{ isData() }

type TestData struct{ 
	DataCode DataCode
	Contents int } 

func (TD *TestData) isData(){}


func CreateTestData(content int) *TestData {
	return &TestData{DataCode: TestDataCode, Contents: content}
}



