package apis

import (
	"cotacao-go/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// initialize API
func InitializeApi() (models.ApiCurrence, error) {

	err := godotenv.Load()
	if err != nil {
		return models.ApiCurrence{}, fmt.Errorf("error: %w", err)
	}

	token := os.Getenv("TOKEN")
	url := os.Getenv("URL") + token

	if token == "" || url == "" {
		return models.ApiCurrence{}, fmt.Errorf("error: Token or URL not found")
	}

	resp, err := http.Get(url)

	if err != nil {
		return models.ApiCurrence{}, fmt.Errorf("error: %w", err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	var result models.ApiCurrence

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return models.ApiCurrence{}, fmt.Errorf("error: %w", err)
	}

	return result, nil
}
