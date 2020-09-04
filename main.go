package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"

	bearer "github.com/bearer/go-agent"
)

func main() {
	agent := bearer.New(
		os.Getenv("BEARER_SECRET_KEY"),
		bearer.WithEnvironment(os.Getenv("APP_ENV")),
	)
	defer agent.Close()

	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	// Postman Echo
	fmt.Println("-- Sending API Calls to Postman-echo --")
	req, _ := http.NewRequest("GET", "https://postman-echo.com/status/404", nil)
	req.Header.Set("password", "h4x0r")
	req.Header.Set("secret", "secret key")
	client.Do(req)

	req2, _ := http.NewRequest("GET", "https://postman-echo.com/status/403", nil)
	req2.Header.Set("password", "h4x0r")
	req2.Header.Set("secret", "secret key")
	client.Do(req2)

	req3, _ := http.NewRequest("GET", "https://postman-echo.com/status/429", nil)
	req3.Header.Set("password", "h4x0r")
	req3.Header.Set("secret", "secret key")
	client.Do(req3)

	req4, _ := http.NewRequest("GET", "https://postman-echo.com/status/501?email=pii@example.com", nil)
	req4.Header.Set("password", "h4x0r")
	req4.Header.Set("secret", "secret key")
	client.Do(req4)

	body5, _ := json.Marshal(map[string]string{
		"foo": "Bar",
	})
	req5, _ := http.NewRequest("POST", "https://postman-echo.com/post", bytes.NewBuffer(body5))
	req5.Header.Set("Content-Type", "application/json")
	client.Do(req5)

	client.Get("https://postman-echo.com/status/200")

	// NASA API
	fmt.Println("-- Sending API Calls to NASA API --")
	client.Get("https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY")

	// Star Wars API
	fmt.Println("-- Sending API Calls to SWAPI --")
	client.Get("https://swapi.dev/api/people/1/")
	client.Get("https://swapi.dev/api/people/9/")
	client.Get("https://swapi.dev/api/starships/9/")

	// Foo Bar -> Non Existing API
	fmt.Println("-- Sending API Calls to non existing API --")
	client.Get("https://foo.bar/status/200")
}
