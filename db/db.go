package db

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/supreme-potato/config"
)

type db struct {
	dyna *dynamodb.DynamoDB
}

var (
	d    *db
	once sync.Once
)

func DB() *db {
	once.Do(func() {
		sess := session.Must(session.NewSession())
		svc := dynamodb.New(sess, aws.NewConfig().WithEndpoint(config.Config().DynamodbEndpoint).WithRegion(config.Config().DynamodbRegion))
		d = &db{
			dyna: svc,
		}
	})
	return d
}

func (d *db) CreateTable(name string) {
	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("name"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("name"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(config.Config().DynamodbReadCapacityUnits),
			WriteCapacityUnits: aws.Int64(config.Config().DynamodbWriteCapacityUnits),
		},
		TableName: aws.String(name),
	}

	resp, err := d.dyna.CreateTable(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func (d *db) ListTables() {
	params := &dynamodb.ListTablesInput{
	// ExclusiveStartTableName: aws.String("TableName"),
	// Limit: aws.Int64(5),
	}
	resp, err := d.dyna.ListTables(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func (d *db) Set(key, data string) {
}

func (d *db) Get(key string) (string, error) {
	return "", nil
}
