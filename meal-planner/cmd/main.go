package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Meal struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type ApiResponse struct {
	Success bool `json:"success"`
	Status  int  `json:"status"`
	Data    struct {
		Page       int    `json:"page"`
		TotalPages int    `json:"total_pages"`
		Meals      []Meal `json:"data"`
	} `json:"data"`
}

func GetMeal() string {
	resp, err := http.Get("https://studio.api.f45training.com/v1/strapi/recipe?page=0&per_page=1")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data ApiResponse
	json.Unmarshal(responseData, &data)

	return fmt.Sprintf("%+v\n", data)
}

func main() {
	fmt.Println("Ran")
	fmt.Print(GetMeal())
}
