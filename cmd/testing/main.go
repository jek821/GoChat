package main

import (
	"GoChat/pkg"
	"fmt"
)

func main() {
	EncodeDecodeTest()
}

func EncodeDecodeTest() {
	pkg.Init()
	dataForTest := 5
	testData :=	pkg.CreateTestData(dataForTest)
	testTransmission :=	pkg.CreateTransmission(testData.DataCode, testData)
	encodedTransmission, err := testTransmission.EncodeTransmission()
	if err != nil {
		fmt.Println("Failed Encoding Test")
	}
	decodedTransmission, err := pkg.DecodeTransmission(encodedTransmission)
	if err != nil {
		fmt.Println("Failed Decoding Test")
	}
	switch decodedTransmission.Code {
	case (pkg.TestDataCode):
		TransmissionData := decodedTransmission.Data.(*pkg.TestData)
		if (TransmissionData.Contents == dataForTest){
		fmt.Println("Encode Decode Test Passed!")
	}
	default:
	fmt.Println("Unknown Data type used in Encode Decode Testing")
	}
}
