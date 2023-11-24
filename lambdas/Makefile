.PHONY: clean build

all: clean build

clean:
	rm -f bootstrap myFunction.zip

build:
	# Build bootstrap
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go

	# Zip it
	zip myFunction.zip bootstrap
