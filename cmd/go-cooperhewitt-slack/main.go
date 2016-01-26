package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"log"
	"fmt"
	"github.com/cooperhewitt/go-cooperhewitt-api"
)

func hello(w http.ResponseWriter, r *http.Request) {

	token := os.Getenv("TOKEN")

	if token == "" {
		log.Fatal("$TOKEN must be set")
	}

	client := api.OAuth2Client(token)

	method := "cooperhewitt.labs.whatWouldMicahSay"
	args := url.Values{}

	rsp, err := client.ExecuteMethod(method, &args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to call %s, because '%s'\n", method, err)
		os.Exit(1)
	}

	_, api_err := rsp.Ok()

	if api_err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute %s, because '%s'\n", method, api_err.Message)
		os.Exit(1)
	}

	body := rsp.Body()

	var says string
	says, _ = body.Path("micah.says").Data().(string)

	io.WriteString(w, says)
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", hello)
	http.ListenAndServe(":" + port, nil)
}
