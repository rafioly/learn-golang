package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	GetRoot(res, req)

	if res.Body.String() != "You're on ROOT" {
		t.Error("Fail! It should tell me that I'm on ROOT, got", res.Body.String(), "instead")
	}
}

func TestListAllHandlerEmpty(t *testing.T) {
	req, _ := http.NewRequest("GET", "/list_all", nil)
	res := httptest.NewRecorder()

	ListAll(res, req)

	if res.Body.String() != "{}" {
		t.Error("Fail! The dict should still empty, got", res.Body.String(), "instead")
	}
}

func TestAddItemHandler(t *testing.T) {
	req, _ := http.NewRequest("POST", "/add_item?name=mangga&price=7000", nil)
	res := httptest.NewRecorder()

	AddItem(res, req)

	if _, ok := supermarket["mangga"]; !ok {
		t.Error("Fail! Item mangga isn't added")
	}

	if res.Body.String() != "Item named mangga with price 7000 has been successfully added" {
		t.Error("Fail! It should show Item named mangga with price 7000 has been successfully added, got", res.Body.String(), "instead")
	}
}

func TestListAllHandlerExist(t *testing.T) {
	req, _ := http.NewRequest("GET", "/list_all", nil)
	res := httptest.NewRecorder()

	ListAll(res, req)

	if res.Body.String() != "{\"mangga\":7000}\n" {
		t.Error("Fail! It should show {\"mangga\":7000}, got", res.Body.String(), "instead")
	}
}

func TestDeleteItemHandler(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/?name=mangga", nil)
	res := httptest.NewRecorder()

	DeleteItem(res, req)

	if _, ok := supermarket["mangga"]; ok {
		t.Error("Fail!")
	}

	if res.Body.String() != "Item named mangga has been successfully removed" {
		t.Error("Fail!")
	}
}
