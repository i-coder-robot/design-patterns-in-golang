package Builder

import (
	"fmt"
	"testing"
)

func TestSender_BuildMessage(t *testing.T) {
	json:=&JSONMessageBuilder{}
	xml:=&XMLMessageBuilder{}
	sender:= Sender{}

	sender.SetBuilder(json)
	jsonMsg,err:= sender.BuildMessage("Golang","Design Patter")
	if err != nil {
		t.Fatalf("BuildMesage(): JSON: no se esperaba error, se recibió: %v", err)
	}
	fmt.Println(string(jsonMsg.Body))

	sender.SetBuilder(xml)
	xmlMsg,err := sender.BuildMessage("Golang","Design Patter")
	if err != nil {
		t.Fatalf("BuildMesage(): XML: no se esperaba error, se recibió: %v", err)
	}
	fmt.Println(string(xmlMsg.Body))
}