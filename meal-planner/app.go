package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// // Greet returns a greeting for the given name
// func (a *App) Greet(name string) string {
// 	return fmt.Sprintf("Hello %s, It's show time!", name)
// }

// Greet returns a greeting for the given name
func (a *App) Data() string {

	resp, err := http.Get("https://studio.api.f45training.com/v1/strapi/recipe?page=0&per_page=10")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	var str string
	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Printf("%v", i)
		str = str + scanner.Text()
	}

	return str
}

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

func (a *App) GetMeal() string {
	resp, err := http.Get("https://studio.api.f45training.com/v1/strapi/recipe?page=0&per_page=10")

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
