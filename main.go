package main

import "github.com/aws/aws-lambda-go/lambda"

// Message - the prefix and list of words to try construct a greeting with
type Message struct {
	Prefix string   `json:"prefix"`
	Slang  []string `json:"slang"`
}

// Response - a list of all the greetings made
type Response struct {
	Greeting []string `json:"greeting"`
}

func abdulHandler(h Message) (r Response, Error error) {

	for _, word := range h.Slang {
		r.Greeting = append(r.Greeting, h.Prefix+" "+word)
	}

	return r, nil
}

func main() {
	lambda.Start(abdulHandler)
}
