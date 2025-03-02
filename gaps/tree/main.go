package main

import (
	"context"
	"encoding/json"

	"github.com/asoliman1/experiments/gaps/internal/pkg/nodes"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/kinesis"
	env "github.com/caarlos0/env/v11"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Environment struct {
	Table  string `env:"TABLE" envDefault:"nodes"`
	Stream string `env:"STREAM" envDefault:"nodes"`
	Region string `env:"AWS_REGION" envDefault:"us-east-2"`
	N      int    `env:"CUTOFF" envDefault:"5"`
}

func main() {
	var e Environment
	if err := env.Parse(&e); err != nil {
		panic(err)
	}

	sess := session.Must(session.NewSession(new(aws.Config).WithRegion(e.Region)))
	krakenClient := kinesis.New(sess)
	dynamoClient := dynamodb.New(sess)

	lambda.Start(func(
		ctx context.Context,
		event events.KinesisEvent,
	) error {
		for _, record := range event.Records {
			var node nodes.Node
			if err := json.Unmarshal(record.Kinesis.Data, &node); err != nil {
				log.Fatal("invalid node record")
			}
			writeToDynamoDB(dynamoClient, e.Table, struct {
				Sequence string `dynamodbav:"sequence"`
				Level    int    `dynamodbav:"level"`
				Value    int    `dynamodbav:"value"`
			}{
				Sequence: node.Sequence,
				Level:    node.Level,
				Value:    node.Value,
			})
			if node.Level == e.N {
				continue
			}

			if leftNode, ok := node.Left(); ok {
				pushToKinesis(krakenClient, e.Stream, leftNode)
			}

			if rightNode, ok := node.Right(); ok {
				pushToKinesis(krakenClient, e.Stream, rightNode)
			}

		}
		return nil
	})
}

func writeToDynamoDB[T any](client *dynamodb.DynamoDB, tableName string, document T) {
	item, err := dynamodbattribute.MarshalMap(document)
	if err != nil {
		log.Fatal("Error marshaling item:", err)
	}
	if _, err := client.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	}); err != nil {
		log.WithFields(log.Fields{
			"item": item,
		}).WithError(err).Fatal("Failed to write to DynamoDB")
	}
}

func pushToKinesis[T any](client *kinesis.Kinesis, streamName string, document T) {
	paritionKey := uuid.New().String()
	data, err := json.Marshal(document)
	if err != nil {
		log.Fatal("invalid node record")
	}
	if _, err := client.PutRecord(&kinesis.PutRecordInput{
		StreamName:   aws.String(streamName),
		Data:         data,
		PartitionKey: aws.String(paritionKey),
	}); err != nil {
		log.WithError(err).Fatal("Failed to put record")
	}
}
