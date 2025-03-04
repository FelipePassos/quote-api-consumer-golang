package generators

import (
	"cotacao-go/models"
	"fmt"
	"os"
)

// generate document files
func GenereatorDocFile(response models.ApiCurrence) error {

	file, err := os.Create("exports/info.pdf")

	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	defer file.Close()

	text := fmt.Sprintf("1 Dollar is equals %.2f BRL", response.USD_BRL.Price)

	_, err = file.WriteString(text)

	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	return nil
}
