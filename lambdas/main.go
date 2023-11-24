package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"log"
	"reflect"
	"strings"

	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ulikunitz/xz"
)

func getType(value interface{}) string {
	return reflect.TypeOf(value).String()
}

func compressImage(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	input := request.Body

	trimmedTheFront := input[strings.IndexByte(input, ',')+1:]
	trimmedTheEnd := trimmedTheFront[:len(trimmedTheFront)-3]

	decodedBytes, err := base64.StdEncoding.DecodeString(trimmedTheEnd)
	if err != nil {
		log.Println("Error decoding Base64: ", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error decoding Base64"}, nil
	}

	decodedImageData, _, err := image.Decode(bytes.NewReader((decodedBytes)))
	if err != nil {
		log.Println("Error decoding image:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error decoding image"}, nil
	}

	var compressedImageBuffer bytes.Buffer
	writer, err := xz.NewWriter(&compressedImageBuffer)
	if err != nil {
		log.Println("Error creating xz writer:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error creating xz writer"}, nil
	}

	err = jpeg.Encode(writer, decodedImageData, nil)
	if err != nil {
		log.Println("Error compressing image:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error compressing image"}, nil
	}

	err = writer.Close()

	if err != nil {
		log.Println("Error closing xz writer:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error closing xz writer"}, nil
	}

	compressedImageData := compressedImageBuffer.Bytes()

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(compressedImageData),
	}, nil
}

func main() {
	lambda.Start(compressImage)
}
