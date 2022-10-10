package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	text2datago "github.com/henomis/text2data-go"
	"github.com/henomis/text2data-go/pkg/request"
)

const PrivateKey = "YOUR_PRIVATE_KEY"
const Secret = "YOUR_SECRET"

func main() {

	text2dataClient := text2datago.New(
		text2datago.Text2DataAPIEndpointV3,
		PrivateKey,
		Secret,
		10*time.Second,
	)

	requestAnalyze := &request.Request{}
	requestAnalyze.DocumentText = "Excellent location, opposite a very large mall with wide variety of shops, restaurants and more."

	response, err := text2dataClient.Analyze(requestAnalyze)
	if err != nil {
		log.Fatalf("error while performing analysis: %v", err)
	}

	if !response.Status.IsSuccess() {
		log.Fatalf("error: %s", response.ErrorMessage.Error())
	}

	bytes, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(bytes))
}
