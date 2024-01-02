package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"Id"`
	Title     string `json:"Title"`
	Completed bool   `json:"Completed"`
}

// GET ISLEMI
func Demo1() {
	response, err := http.Get("http://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err) //Hata var ise hata mesaji
	}
	defer response.Body.Close() //islem bitiminde kapatma.

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)
}

// POST ISLEMI
func Demo2() {
	todo := Todo{1, 2, "Alisveris yapilacak", false}
	jsonTodo, err := json.Marshal(todo) //jaon formatina donusturme islemi

	response, err := http.Post("http://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse)
	fmt.Println(todoResponse)
}
