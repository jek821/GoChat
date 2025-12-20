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
	testData :=	pkg.CreateTestData(5)
	testTransmission :=	pkg.CreateTransmission(testData.DataCode, testData)
	encodedTransmission, err := testTransmission.EncodeTransmission()
	if err != nil {
		fmt.Println("Failed to encode transmission")
	}
	decodedTransmission, err := pkg.DecodeTransmission(encodedTransmission)
	if err != nil {
		fmt.Println("failed to decode transmission")
	}
	switch decodedTransmission.Code {
	case (pkg.TestDataCode):
		TransmissionData := decodedTransmission.Data.(*pkg.TestData)
		fmt.Println(TransmissionData.Contents)
		
	}
}
