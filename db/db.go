package db

import (
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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
		svc := dynamodb.New(sess, aws.NewConfig().WithEndpoint("http://localhost:8000").WithRegion("us-west-2"))
		d = &db{
			dyna: svc,
		}
	})
	return d
}

func (d *db) ListTabls() {
	params := &dynamodb.ListTablesInput{
		ExclusiveStartTableName: aws.String("TableName"),
		Limit: aws.Int64(1),
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
