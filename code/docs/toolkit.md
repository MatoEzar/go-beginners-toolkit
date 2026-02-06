# Prompt-Powered Kickstart: Building a Beginner's Toolkit for Go (Golang)

## 1. Title & Objective
**Technology:** Go (Golang)  
**Why chosen:** Go has clean, readable syntax, excellent performance, and a strong standard library — ideal for beginners to quickly build reliable CLI tools without complex dependencies. It's widely used in modern infrastructure (e.g., Docker, Kubernetes).  
**Objective:** Create a beginner-friendly guide + runnable CLI that fetches and prints random programming jokes using AI-assisted learning.

## 2. Quick Summary of the Technology
Go is an open-source, statically typed, compiled language developed by Google. It emphasizes simplicity, efficiency, and concurrency.  
**Where used:** Backend services, cloud tools, CLI applications, DevOps.  
**Real-world example:** Docker uses Go for its speed and cross-platform builds.

## 3. System Requirements
- Operating System: Windows 10/11 (or macOS/Linux)  
- Go compiler: Version 1.21 or higher (download from https://go.dev/dl/)  
- Text editor: VS Code (with Go extension for better support)  
- Internet connection (for the joke API)  
- No additional packages required (uses only Go standard library)

## 4. Installation & Setup Instructions
1. Download and install Go from https://go.dev/dl/ (select Windows installer; ensure "Add Go to PATH" is checked).  
2. Verify installation: Open VS Code → Terminal → New Terminal → run `go version`.  
3. Create a project folder (e.g., `go-toolkit`).  
4. Inside it, create a subfolder called `code`.  
5. In `code`, create a file named `main.go`.  
6. Paste the code from section 5.  
7. In VS Code terminal: Navigate with `cd code`, then run `go run main.go`.

## 5. Minimal Working Example
```go
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
	fmt.Println("Fetching a random programming joke for you...")

	resp, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching joke: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		os.Exit(1)
	}

	var joke Joke
	err = json.Unmarshal(body, &joke)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing joke JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nJoke:")
	fmt.Println(joke.Setup)
	fmt.Println(joke.Punchline)
}