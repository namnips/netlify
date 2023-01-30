package main

import (
	"encoding/json"
	"fmt"

	mediumautopost "github.com/askcloudarchitech/mediumautopost"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)
type RequestBody struct {
	Payload Payload `json:"payload"`
}

type Payload struct {
	Context string `json:"context"`
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	
	

	requestBody:=RequestBody{}
	json.Unmarshal([]byte(request.Body),&requestBody)

	if requestBody.Payload.Context == "production" {
		mediumautopost.Do("")
	} else {
		fmt.Println("context "+ requestBody.Payload.Context +" detected, skipping")
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		// Headers:           map[string]string{"Content-Type": "text/plain"},
		// MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
		Body:       "Hello, World",
		IsBase64Encoded:   false,
	  }, nil
}

func main() {
  lambda.Start(handler)
}
