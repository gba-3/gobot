GOCMD=go
GOBUILD=${GOCMD} build
BINARY_NAME=gobot
BUCKET_NAME=gobot-application

build:
	GOOS=linux GOARCH=amd64 ${GOBUILD} -o ${BINARY_NAME} ./example/*.go
sam-package:
	sam package  --template-file template.yaml --s3-bucket ${BUCKET_NAME} --output-template-file packaged.yaml
sam-deploy:
	sam deploy --template-file ./packaged.yaml --stack-name gobot-app --region ap-northeast-1 --capabilities CAPABILITY_IAM