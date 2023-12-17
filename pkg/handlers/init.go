package handlers

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/[username/repository-name]/pkg/service"
)

var (
	svc dynamodbiface.DynamoDBAPI
)

func init() {
	svc = service.DynamoDBInit()
}
