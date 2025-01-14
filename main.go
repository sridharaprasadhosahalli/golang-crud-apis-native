package main

import(
	"net/http"
	"encoding/json"
	"strings"
)

type item struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Type string  `json:"type"`
    Price  float64 `json:"price"`
}

// items slice containing items available in a grocery store
var items = []item{
    {ID: "1", Name: "Banana", Type: "Fruit", Price: 40},
    {ID: "2", Name: "Tomoto", Type: "Vegetable", Price: 30},
    {ID: "3", Name: "Cocumber", Type: "Vegetable", Price: 20},
}

func main(){
	http.HandleFunc("/items",getorpostItems)
	http.HandleFunc("/items/",getorputordeleteItems)
	http.ListenAndServe(":8080",nil)
}


// getItems responds with the list of all getItems present in a grocery store as JSON.
func getorpostItems(w http.ResponseWriter,r *http.Request) {
	if(r.Method == "GET"){
		getItems(w,r)
		return
	}else{
		addItem(w,r)
		return
	}
 }

 // getItems responds with the list of all getItems present in a grocery store as JSON.
func getorputordeleteItems(w http.ResponseWriter,r *http.Request) {
	if(r.Method == "GET"){
		getItemByID(w,r)
		return
	}else if(r.Method == "PUT"){
		updateItem(w,r)
		return
	}else{
		deleteItem(w,r)
		return
	}
 }

// getItems responds with the list of all getItems present in a grocery store as JSON.
func getItems(w http.ResponseWriter,r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

// getItemByID searches the item with the same id in grocery store and sends response as JSON.
func getItemByID(w http.ResponseWriter,r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path,"/items/")

    // Loop through the list of items in grocery store, looking for
    // an item whose ID value matches the parameter.
    for _, a := range items {
        if a.ID == id {
			w.Header().Set("Content-Type", "application/json")
	        w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(a)
            return
        }
    }
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("{message:item not found in grocery store.}")
}

// adddItem adds the item from JSON received into grocery store inventory.
func addItem(w http.ResponseWriter,r *http.Request) {
    var newItem item

    err := json.NewDecoder(r.Body).Decode(&newItem)
    if err != nil {
    	http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Add the new item to the grocery store.
    items = append(items, newItem)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newItem)
}

//updateItem searches the item with the same id in grocery store and updates it.
func updateItem(w http.ResponseWriter,r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path,"/items/")
	var newItem item
	
	err := json.NewDecoder(r.Body).Decode(&newItem)
    if err != nil {
    	http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	for i, a := range items {
		if a.ID == id {
			newItem.ID = id // set the id of newItem
			items[i] = newItem
			w.Header().Set("Content-Type", "application/json")
	        w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(a)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("{message:item not found in grocery store.}")
}

//deleteItem delets the item from the grocery store .
func deleteItem(w http.ResponseWriter,r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path,"/items/")
	// Loop through the list of items in grocery store, looking for
    // an item whose ID value matches the parameter and delete it
	for i, a := range items {
        if a.ID == id {
            items = append(items[:i], items[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
	        w.WriteHeader(http.StatusOK)
        }
    }
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("{message:item not found in grocery store.}")
}



