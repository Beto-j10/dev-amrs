package ex2

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Pokemon is a struct that contains the name of the pokemon
type Pokemon struct {
	Name string `json:"name"`
}

// Ex2 prints the name of the pokemon
func Ex2() error {

	fmt.Println("Ejercicio 2")
	fmt.Println("Ingrese id de pokemon:")

	var id int
	fmt.Scanln(&id)

	pokemon, err := getPokemon(id)
	if err != nil {
		log.Printf("Error getting pokemon: %v", err)
		return err
	}

	fmt.Printf("\nPokemon: %s\n\n", pokemon)
	return nil

}

// getPokemon makes a request to the pokeapi and returns the pokemon name
func getPokemon(id int) (string, error) {

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-form/%d/", id) // other way: strings.Builder, strings.Join, plus operator
	client := &http.Client{Timeout: 30 * time.Second}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return "", err
	}

	request.Header.Set("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Error response status: %v", response.StatusCode)
		return "", fmt.Errorf("error response status: %v", response.StatusCode)
	}

	pokemon := Pokemon{}

	err = json.NewDecoder(response.Body).Decode(&pokemon)
	if err != nil {
		return "", err
	}

	return pokemon.Name, nil
}
