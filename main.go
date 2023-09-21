package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todos struct {
	UserID    int64  `json:"userId"`
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	//api call

	url := "https://jsonplaceholder.typicode.com/todos/1/"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {

		//allow only field in stuct
		todoItem := Todos{}
		json.NewDecoder(res.Body).Decode(&todoItem)
		//convert back to json
		todo, err := json.MarshalIndent(todoItem, "", "\t")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(todo))

		// //dont allow unknownfields
		// todoItem := Todos{}
		// decoder := json.NewDecoder(res.Body)
		// decoder.DisallowUnknownFields()
		// if err := decoder.Decode(&todoItem); err != nil {
		// 	log.Fatal("Error:", err)
		// }
		//convert back to json
		// todo, err := json.Marshal(todoItem)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(string(todo))

	}

}
