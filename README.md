# Text2Data SDK for Go


[![Build Status](https://github.com/henomis/text2data-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/henomis/text2data-go/actions/workflows/test.yml?query=branch%3Amain) [![GoDoc](https://godoc.org/github.com/henomis/text2data-go?status.svg)](https://godoc.org/github.com/henomis/text2data-go) [![Go Report Card](https://goreportcard.com/badge/github.com/henomis/text2data-go)](https://goreportcard.com/report/github.com/henomis/text2data-go) [![GitHub release](https://img.shields.io/github/release/henomis/text2data-go.svg)](https://github.com/henomis/text2data-go/releases)

This is Text2Data's **unofficial** Go client, designed to enable you to use Text2Data's services easily from your own applications.

## Text2Data

Text2Data is a cloud-based text analytics service that through APIs allows you extract informations from a text content.

## SDK versions

This SDK is compatible with v3 of the Text2Data API.

## Getting started

### Installation

You can load text2data-go into your project by using:
```
go get github.com/henomis/text2data-go
```


### Configuration

The only things you need to start using Text2Data's APIs are `PrivateKey` and `Secret`. 


### Usage

Please refer to the [examples folder](examples/) to see how to use the SDK.

Here below a simple usage example:

```go
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
```