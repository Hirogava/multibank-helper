package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Hirogava/multibank-helper/internal/config/logger"
	"github.com/segmentio/kafka-go"
)

type KafkaManager struct {
    Brokers []string
    Writer  *kafka.Writer
    Reader  *kafka.Reader
}

func NewKafkaManager(brokers []string, topic, groupID string) *KafkaManager {
	logger.Logger.Infof("Create new KafkaManager with brokers: %v, topic: %s, groupID: %s", brokers, topic, groupID)

    writer := &kafka.Writer{
        Addr:         kafka.TCP(brokers...),
        Topic:        topic,
        Balancer:     &kafka.LeastBytes{},
        RequiredAcks: kafka.RequireAll,
        Async:        false,
    }

    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers:  brokers,
        Topic:    topic,
        GroupID:  groupID,
        MinBytes: 10e3,
        MaxBytes: 10e6,
    })

    return &KafkaManager{
        Brokers: brokers,
        Writer:  writer,
        Reader:  reader,
    }
}

func (km *KafkaManager) SendMessage(ctx context.Context, key string, value interface{}) error {
    data, err := json.Marshal(value)
    if err != nil {
        return err
    }

    msg := kafka.Message{
        Key:   []byte(key),
        Value: []byte(data),
        Time:  time.Now(),
    }

    if err := km.Writer.WriteMessages(ctx, msg); err != nil {
        return err
    }

    logger.Logger.Infof("Sent message: key=%s value=%s", key, value)
    return nil
}

func (km *KafkaManager) ReadMessage(ctx context.Context) error {
    msg, err := km.Reader.ReadMessage(ctx)
    if err != nil {
        return err
    }

    logger.Logger.Infof("Received message: key=%s value=%s", string(msg.Key), string(msg.Value))
    return nil
}

func (km *KafkaManager) ReadMessages(ctx context.Context) {
	logger.Logger.Info("Starting to read messages")

	for {
		_, err := km.Reader.ReadMessage(ctx)
		if err != nil {
			logger.Logger.Errorf("Error reading message: %v", err)
			continue
		}

		// TODO: Handle the message
	}
}

func (km *KafkaManager) Close() error {
    if err := km.Writer.Close(); err != nil {
        return err
    }
    if err := km.Reader.Close(); err != nil {
        return err
    }
    return nil
}
