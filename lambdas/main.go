package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
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

	fmt.Println("0")

	decodedBytes, err := base64.StdEncoding.DecodeString(trimmedTheEnd)
	if err != nil {
		log.Println("Error decoding Base64: ", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error decoding Base64"}, nil
	}

	fmt.Println("1")

	decodedImageData, _, err := image.Decode(bytes.NewReader((decodedBytes)))
	if err != nil {
		log.Println("Error decoding image:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error decoding image"}, nil
	}

	fmt.Println("2")

	var compressedImageBuffer bytes.Buffer
	writer, err := xz.NewWriter(&compressedImageBuffer)
	if err != nil {
		log.Println("Error creating xz writer:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error creating xz writer"}, nil
	}

	fmt.Println("4")

	err = jpeg.Encode(writer, decodedImageData, nil)
	if err != nil {
		log.Println("Error compressing image:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error compressing image"}, nil
	}

	fmt.Println("5")

	err = writer.Close()

	fmt.Println("6")

	if err != nil {
		log.Println("Error closing xz writer:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error closing xz writer"}, nil
	}

	fmt.Println("7")
	// convert the compressed image buffer to a base64-encoded string
	compressedImageData := compressedImageBuffer.Bytes()

	fmt.Println("8")
	// save the compressed image to a file or upload it to S3 here if needed

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(compressedImageData),
	}, nil
}

func main() {
	lambda.Start(compressImage)
}
