package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Nama"`
	Umur     int
}

func main() {
	var jsonString = `{"Nama": "Elang Muhammad FIkhri", "Umur": 21}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Nama :", data.FullName)
	fmt.Println("Umur  :", data.Umur)
}
