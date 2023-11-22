package main

import (
	"bytes"
	"image"
	"log"

	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ulikunitz/xz"
)

func compressImage(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	imageData := request.Body
	log.Println("Received image data:", imageData)

	img, _, err := image.Decode(bytes.NewReader([]byte(imageData)))
	if err != nil {
		log.Println("Error decoding image:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error decoding image"}, nil
	}

	// Compress the image using xz compression
	var compressedImageBuffer bytes.Buffer
	writer, err := xz.NewWriter(&compressedImageBuffer)
	if err != nil {
		log.Println("Error creating xz writer:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error creating xz writer"}, nil
	}

	err = jpeg.Encode(writer, img, nil)
	if err != nil {
		log.Println("Error compressing image:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error compressing image"}, nil
	}

	err = writer.Close()
	if err != nil {
		log.Println("Error closing xz writer:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error closing xz writer"}, nil
	}

	// Convert the compressed image buffer to a base64-encoded string
	compressedImageData := compressedImageBuffer.Bytes()

	// You can save the compressed image to a file or upload it to S3 here if needed

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(compressedImageData),
	}, nil
}

func main() {
	lambda.Start(compressImage)
}
