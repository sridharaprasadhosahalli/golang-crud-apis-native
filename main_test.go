package main

import(
    "net/http"
    "net/http/httptest"
    "testing"
	"encoding/json"
)


/*func TestGetItemsHandler(t *testing.T) {

    req, _ := http.NewRequest("GET", "/items", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    var items []item
    json.Unmarshal(w.Body.Bytes(), &items)

}*/

/*func TestGetItemsHandler(t *testing.T){
    res, err := http.Get("http://localhost:8080/items")
    if err != nil {
		log.Fatal(err)
	}
    body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}*/

func TestGetItemsHandler(t *testing.T){

	req := httptest.NewRequest(http.MethodGet, "/items", nil)
    w := httptest.NewRecorder()
	getorpostItems(w,req)
    var items []item
    json.Unmarshal(w.Body.Bytes(), &items)
	if(w.Code != http.StatusOK || len(items) < 3){
		t.Error("Test Failed ", w.Code)
	}



}