package handlers

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/Robson704/Projekttt/pkg/service"
)

var (
	svc dynamodbiface.DynamoDBAPI
)

func init() {
	svc = service.DynamoDBInit()
}
