package main

import(
    "net/http"
    "net/http/httptest"
    "testing"
	"encoding/json"
	"bytes"
)

func TestGetItemsHandler(t *testing.T){

	req := httptest.NewRequest("GET", "/items", nil)
    w := httptest.NewRecorder()
	getorpostItems(w,req)
    var items []item
    json.Unmarshal(w.Body.Bytes(), &items)
	if(w.Code != http.StatusOK || len(items) < 3){
		t.Error("Test Failed ", w.Code)
	}
}

func TestGetItemByIdHandler(t *testing.T){

	req := httptest.NewRequest("GET", "/items/1", nil)
    w := httptest.NewRecorder()
	getorputordeleteItems(w,req)
    var item []item
    json.Unmarshal(w.Body.Bytes(), &item)
	if(w.Code != http.StatusOK || len(items) < 0){
		t.Error("Test Failed ", w.Code)
	}
}

func TestDeleteItemByIdHandler(t *testing.T){

	req := httptest.NewRequest("DELETE", "/items/3", nil)
    w := httptest.NewRecorder()
	getorputordeleteItems(w,req)
	if(w.Code != http.StatusOK || len(items) != 2){
		t.Error("Test Failed ", w.Code)
	}
}

func TestAddItemHandler(t *testing.T){
	itemm := item{
        ID: "4",
        Name: "Orange",
        Type: "Fruit",
        Price: 60,
    }
    jsonValue, _ := json.Marshal(itemm)
	req := httptest.NewRequest("POST", "/items", bytes.NewBuffer(jsonValue))
    w := httptest.NewRecorder()
	getorpostItems(w,req)
	if(w.Code != http.StatusCreated || len(items) < 0){
		t.Error("Test Failed ", w.Code)
	}
}

func TestUpdateItemHandler(t *testing.T){
	itemm := item{
        ID: "1",
        Name: "Orange",
        Type: "Fruit",
        Price: 60,
    }
    jsonValue, _ := json.Marshal(itemm)
	req:= httptest.NewRequest("PUT", "/items/1", bytes.NewBuffer(jsonValue))
    w := httptest.NewRecorder()
	getorpostItems(w,req)
	if(w.Code != http.StatusCreated || len(items) < 0 ){
		t.Error("Test Failed ", w.Code)
	}
}