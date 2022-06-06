package _select

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func NewKafkaReader(brokers []string, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB

	})
}

func consumeKafka(ctx context.Context, reader *kafka.Reader, queue []kafka.Message) {
	for count := 0; count < 10; count++ {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("read message data : %s", string(m.Value[:]))
		queue = append(queue, m)
	}
}

func closeKafka(reader *kafka.Reader) {
	fmt.Println("start close kafka")
	reader.Close()
	fmt.Println("close kafka successfully")
}

func NewTimer(reader *kafka.Reader) {
	//cxt -> main context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	now := time.Now()

	var queue []kafka.Message
	go func(ctx context.Context) {
		consumeKafka(ctx, reader, queue)
	}(ctx)
	defer closeKafka(reader)

	log.Printf("%d", time.Since(now))
}
