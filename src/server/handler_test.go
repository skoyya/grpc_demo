package server

import (
	"testing"
	"api"
)

func TestSayHello(t *testing.T) {
	s := Server{}
	resp , err := s.SayHello(nil,&api.RequestMessage{Message :"Dummy"})
	if err != nil {
		t.Errorf("Server SayHello failed" ,err)
		return
	}

	if resp == nil { 
		t.Error("Server SayHello failed, nil") 
		return
	}

	want := "No.... no handshake. Corona time"
	if resp.Message != want {
		t.Errorf("Server SayHello failed, Expected[%s],got[%s]",want,resp.Message)
	}
}
