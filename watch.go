package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := getEnv("TOPIC", "test")
	partition := getEnv("PARTITION", "0")
	offset := getEnv("OFFSET", "0")
	broker := getEnv("BROKER", "kafka-headless:9092")

	p, err := strconv.Atoi(partition)
	if err != nil {
		fmt.Println("PARTITION must be an integer.")
		os.Exit(1)
	}

	o, err := strconv.ParseInt(offset, 10, 64)
	if err != nil {
		fmt.Println("OFFSET must be an integer.")
		os.Exit(1)
	}

	fmt.Println("Starting Kafka Watch...")

	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     topic,
		Partition: p,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	r.SetOffset(o)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	r.Close()
}

// getEnv gets an environment variable or sets a default if
// one does not exist.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}
