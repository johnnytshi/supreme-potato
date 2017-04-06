package config

import "sync"

type config struct {
	DynamodbEndpoint           string
	DynamodbRegion             string
	DynamodbReadCapacityUnits  int
	DynamodbWriteCapacityUnits int
}

var (
	c    *config
	once sync.Once
)

func Config() *config {
	once.Do(func() {
		c = &config{
			DynamodbEndpoint:           "http://localhost:8000",
			DynamodbRegion:             "us-west-2",
			DynamodbReadCapacityUnits:  1,
			DynamodbWriteCapacityUnits: 1,
		}
	})
	return c
}
