package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Refractoring:Yeniden Yapilandirma

type Products struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Products, error) {
	response, err := http.Get("http://localhost:3000/products")

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []Products
	json.Unmarshal(bodyBytes, &products)
	//fmt.Println(products)
	return products, nil
}

func AddProduct() (Products, error) {
	// product := Products{Id: 4, ProductName: "Telephone", CategoryId: 1, UnitPrice: 6000.0}
	//product := Products{ProductName: "Microphone", CategoryId: 1, UnitPrice: 2000.0}
	//product := Products{Id: 8, ProductName: "Microphone", CategoryId: 1, UnitPrice: 2000.0}
	//product := Products{Id: 2, ProductName: "Microphone2", CategoryId: 1, UnitPrice: 2000.0}
	product := Products{ProductName: "Microphone3", CategoryId: 1, UnitPrice: 2000.0}

	jsonProduct, err := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))

	/*if err != nil {
		fmt.Println(err)
	}
	*/
	if err != nil {
		return Products{}, err
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var productResponse Products
	json.Unmarshal(bodyBytes, &productResponse)

	//fmt.Println("Kaydedildi : ", productResponse)
	return productResponse, nil
}
