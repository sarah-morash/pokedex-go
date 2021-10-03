package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// A Response struct to map the Entire Response
type AllPokemonResponse struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
	ID int `json:"id"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}

func main() {
	// INTRO
	fmt.Println("HELLO! Welcome to my Pokedex 🎉")
	fmt.Println("1: Get info on a specific pokemon")
	fmt.Println("2: List all pokemon by region")
	fmt.Println("3: Exit")
	fmt.Print("Select an option: ")

	// Get input from user by standard input
	var menuOption string
	fmt.Scanln(&menuOption)
	fmt.Println("")

	// Menu option logic 
	// 1 - get a specific pokemon
	// 2 - get all pokemon by region
	// 3 - exit
	if menuOption == "1" {
		getSpecificPokemon()

	} else if menuOption == "2" {
		getAllPokemon()

	} else if menuOption == "3" {
		fmt.Println("Goodbye! 👋")

	} else {
		fmt.Println("Not a menu option")
	}
}

// Option 2: get all pokemon
func getAllPokemon() {
	fmt.Println("Get all Pokemon by region")
	fmt.Print("Please select a region (ex: kanto/hoenn/galar): ")

	// Get input from user by standard input and make it lower case
	var input string
	fmt.Scanln(&input)
	fmt.Println("")
	input = strings.ToLower(input)

	// Try to match with one of the regions we know that exist
	matched, err := regexp.Match(`kanto|hoenn|galar`, []byte(input))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	if (matched) {
		// Try to call API
		pokemonAPI := fmt.Sprintf("https://pokeapi.co/api/v2/pokedex/%s/", input)
		response, err := http.Get(pokemonAPI)

		// Error handling
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		// Read response and store in responseData or get error and handle
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Unpack object into JSON format and print details
		var responseObject AllPokemonResponse
		json.Unmarshal(responseData, &responseObject)

		fmt.Println(responseObject.Name)
		fmt.Println(len(responseObject.Pokemon))

		for i := 0; i < len(responseObject.Pokemon); i++ {
			fmt.Println(responseObject.Pokemon[i].Species.Name)
		}
	} else {
		fmt.Println("The requested region was not found.")
	}
}

func getSpecificPokemon() {
	fmt.Println("Get a specific Pokemon")
	fmt.Print("Please type in a pokemon name or ID: ")

	// Get input from user by standard input and make it lower case
	var input string
	fmt.Scanln(&input)
	fmt.Println("")
	input = strings.ToLower(input)

	// Try to call API
	pokemonAPI := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", input)
	response, err := http.Get(pokemonAPI)

	// Error handling
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Read response and store in responseData or get error and handle
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unpack object into JSON format and print details
	var responseObject PokemonSpecies
	json.Unmarshal(responseData, &responseObject)

	res := fmt.Sprintf("Name: %s\nID: %d\nBase Experience: %d\nHeight: %d\nWeight: %d", responseObject.Name, responseObject.ID, responseObject.BaseExperience, responseObject.Height, responseObject.Weight)
	fmt.Println(res)
}