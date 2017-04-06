// pulls the latest Specs repo and parse the podspec.json files
package main

import (
	"fmt"
	"os"

	"github.com/supreme-potato/db"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	fmt.Println("It's better to be a pirate than join the navy.")
	(&cli.App{}).Run(os.Args)

	db.DB().ListTabls()

	// sess := session.Must(session.NewSession())
	//
	// var svc *DynamoDB
	// svc = dynamodb.New(sess, aws.NewConfig().WithEndpoint("http://localhost:8000").WithRegion("us-west-2"))
	//
	// params := &dynamodb.ListTablesInput{
	// 	ExclusiveStartTableName: aws.String("TableName"),
	// 	Limit: aws.Int64(1),
	// }
	// resp, err := svc.ListTables(params)
	//
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	//
	// fmt.Println(resp)

	// list folders

	// parse json

	// dump into db

	// repeat
}
