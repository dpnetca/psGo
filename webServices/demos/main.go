package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar called"))
}

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

var productList []Product

func init() {
	productsJson := `[
		{
		  "productId": 1,
		  "manufacturer": "Johns-Jenkins",
		  "sku": "p5z343vdS",
		  "upc": "939581000000",
		  "pricePerUnit": "497.45",
		  "quantityOnHand": 9703,
		  "productName": "sticky note"
		},
		{
		  "productId": 2,
		  "manufacturer": "Hessel, Schimmel and Feeney",
		  "sku": "i7v300kmx",
		  "upc": "740979000000",
		  "pricePerUnit": "282.29",
		  "quantityOnHand": 9217,
		  "productName": "leg warmers"
		},
		{
		  "productId": 3,
		  "manufacturer": "Swaniawski, Bartoletti and Bruen",
		  "sku": "q0L657ys7",
		  "upc": "111730000000",
		  "pricePerUnit": "436.26",
		  "quantityOnHand": 5905,
		  "productName": "lamp shade"
		},
	]`
	err := json.Unmarshal([]byte(productsJson), &productList)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	// option 1: handler
	http.Handle("/foo", &fooHandler{Message: "foo called"})
	// option 2 handle func
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":5000", nil)

	product := &Product{
		ProductID:      123,
		Manufacturer:   "Big Box Company",
		Sku:            "4561qHJK",
		Upc:            "414654444566",
		PricePerUnit:   "12.99",
		QuantityOnHand: 28,
		ProductName:    "Gizmo",
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(productJSON))

	jsonProduct := `{
		"productId":123,
		"manufacturer":"Big Box Company",
		"sku":"4561qHJK",
		"upc":"414654444566",
		"pricePerUnit":"12.99",
		"quantityOnHand":28,
		"productName":"Gizmo"
		}`
	newProduct := Product{}
	err = json.Unmarshal([]byte(jsonProduct), &newProduct)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newProduct.ProductName)
}
