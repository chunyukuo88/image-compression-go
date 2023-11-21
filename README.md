# Welcome to your CDK TypeScript project

This is a blank project for CDK development with TypeScript.

The `cdk.json` file tells the CDK Toolkit how to execute your app.

## Useful commands

* `npm run build`   compile typescript to js
* `npm run watch`   watch for changes and compile
* `npm run test`    perform the jest unit tests
* `cdk deploy`      deploy this stack to your default AWS account/region
* `cdk diff`        compare deployed stack with current state
* `cdk synth`       emits the synthesized CloudFormation template


The official docs:
https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html

搞砸了嗎。 沒事，這樣重做：

- delete the original main executable. 
- make changes
- in the lambdas directory, do `GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go`
- `zip myFunction.zip bootstrap` <== Make sure this is specified in the aws-lambda-go-stack.ts file.
- cd ..
- cdk deploy