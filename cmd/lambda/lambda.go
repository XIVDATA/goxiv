package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/xivdata/goxiv"
	"github.com/xivdata/goxiv/model/character"
)

type Request struct {
	ID int64 `json:"id"`
}

func HandleRequest(ctx context.Context, id Request) (character.Character, error) {
	scraper := goxiv.GoXIV{}
	temp := scraper.ScrapeCharacter(id.ID)
	return temp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
