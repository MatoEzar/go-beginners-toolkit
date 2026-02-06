package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func main() {
	fmt.Println("Fetching a random programming joke...")

	resp, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not fetch the joke (%v)\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not read the response (%v)\n", err)
		os.Exit(1)
	}

	var joke Joke
	err = json.Unmarshal(body, &joke)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not understand the joke data (%v)\n", err)
		os.Exit(1)
	}

	fmt.Println("\nHere's your joke:")
	fmt.Println(joke.Setup)
	fmt.Println(joke.Punchline)
}
