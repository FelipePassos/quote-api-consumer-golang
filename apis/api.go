package apis

import (
	"cotacao-go/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// inicializa a API
func InitializeApi() (models.ApiCurrence, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")
	url := os.Getenv("URL") + token

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error body read")
	}

	var result models.ApiCurrence
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal("Error JSON unmarshal")
	}

	return result, nil
}
