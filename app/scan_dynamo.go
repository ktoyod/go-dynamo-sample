/*
https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/gov2/dynamodb/ScanItems/ScanItemsv2.go
*/
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBScanAPI interface {
	Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
}

type Item struct {
	Id   string
	Name string
}

func GetItems(c context.Context, api DynamoDBScanAPI, input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return api.Scan(c, input)
}

func scanDynamoDB(table *string, id *string) []Item {
	filt1 := expression.Name("id").Equal(expression.Value(id))

	expr, err := expression.NewBuilder().WithFilter(filt1).Build()
	if err != nil {
		fmt.Println("Got Error building expression:")
		fmt.Println(err.Error())
		return []Item{}
	}

	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 table,
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	client := dynamodb.NewFromConfig(cfg)

	resp, err := GetItems(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error scanning the table:")
		fmt.Println(err.Error())
		return []Item{}
	}

	items := []Item{}

	err = attributevalue.UnmarshalListOfMaps(resp.Items, &items)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dyanmodb Scan Items, %v", err))
	}

	return items
}
