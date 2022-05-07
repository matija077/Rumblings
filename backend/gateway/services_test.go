package main

import "testing"

func Test_CreateServicesModel(t *testing.T) {
	var serviceModel = createServicesModel()
	frontend, ok := serviceModel.getFrontend("/")
	if !ok {
		t.Error("Frontedn not existing")
	}

	if frontend.service.port != frontend.port {
		t.Error("Frontedn ports not the same")
	}

}
