package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lamdba-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/teris-io/shortid"
)

const (
	TABLE_NAME = "dummy_dynamodb"
	REGION = "us-esat-1"
)

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	ShortURL string `json: "short_url"`
}

type Link struct {
	ShortURL string
	LongURL string
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res := events.APIGatewayProxyResponse{
		Headers: make(map[string]string)
	}

	res.Headers["Access-Control-Allow-Origin"] = "*"
	rb := Request{}
	if err := json.Unmarshal([]byte(request.Body), &rb); err != nil {
		return res, err
	}

	session_, err := session.NewSession(&aws.Config{
		Region: aws.String(REGION)
	})
	if err != nil {
		return res, err
	}

	svc := dynamodb.New(session_)

	shortUrl := shortid.MustGenerate()
	for shortUrl == "shorten" {
		shortUrl = shorid.MustGenerate()
	}
	link := &Link{
		ShortURL: shortUrl,
		LongURL: rb.URL,
	}

	av, err := dynamodbattribute.MarshalMap(link)
	if err != nil {
		return res, err
	}

	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String(TABLE_NAME)
	}
	if _, err := svc.PutItem(input); err != nil {
		return res, err
	}

	response, err := json.Marshal(Response{ShortURL: shortUrl})
	if err != nil {
		return res, err
	}

	res.StatusCode = http.StatusOk
	res.Body = string(response)
	
	return res, nil
}

func main() {
	lamdba.Start(Handler)
}
