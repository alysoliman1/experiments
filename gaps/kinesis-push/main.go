package main

import (
	"encoding/json"
	"log"

	"github.com/asoliman1/experiments/gaps/internal/pkg/nodes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/caarlos0/env/v11"
	"github.com/google/uuid"
)

type Environment struct {
	Stream string `env:"STREAM" envDefault:"nodes"`
	Region string `env:"AWS_REGION" envDefault:"us-east-2"`
}

func main() {
	var e Environment
	if err := env.Parse(&e); err != nil {
		panic(err)
	}

	node := nodes.Node{
		Sequence:        "",
		GapDepth:        100,
		Lengths:         []int{},
		Tails:           []int{},
		LastRecurrences: []int{},
		GapBuckets:      [][]int{},
		Level:           0,
		Value:           -1,
	}

	for range node.GapDepth {
		node.Lengths = append(node.Lengths, 0)
		node.Tails = append(node.Tails, -1)
		node.LastRecurrences = append(node.LastRecurrences, -1)
		node.GapBuckets = append(node.GapBuckets, []int{})
	}

	sess := session.Must(session.NewSession(new(aws.Config).WithRegion(e.Region)))
	krakenClient := kinesis.New(sess)
	paritionKey := uuid.New().String()
	data, err := json.Marshal(node)
	if err != nil {
		log.Fatal("invalid node record")
	}
	res, err := krakenClient.PutRecord(&kinesis.PutRecordInput{
		StreamName:   aws.String(e.Stream),
		Data:         data,
		PartitionKey: aws.String(paritionKey),
	})
	if err != nil {
		log.Fatalf("Failed to put record: %v", err)
	}
	log.Println(aws.StringValue(res.ShardId), aws.StringValue(res.SequenceNumber))
}
