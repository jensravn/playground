package gob

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"

	"cloud.google.com/go/pubsub"
)

type PushRequest struct {
	Message      pubsub.Message `json:"message"`
	Subscription string         `json:"subscription"`
}

func DecodeJSON[E any](r io.Reader) (*E, error) {
	var pr PushRequest
	err := json.NewDecoder(r).Decode(&pr)
	if err != nil {
		return nil, fmt.Errorf("decode push request: %w", err)
	}
	buf := bytes.NewBuffer(pr.Message.Data)
	var entity E
	err = gob.NewDecoder(buf).Decode(&entity)
	if err != nil {
		return nil, fmt.Errorf("decode entity: %w", err)
	}
	return &entity, nil
}

func Decode[E any](binary []byte) (*E, error) {
	buf := bytes.NewBuffer(binary)
	var entity E
	err := gob.NewDecoder(buf).Decode(&entity)
	if err != nil {
		return nil, fmt.Errorf("decode entity: %w", err)
	}
	return &entity, nil
}

func Encode(entity any) ([]byte, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(entity)
	if err != nil {
		return nil, fmt.Errorf("encode entity: %w", err)
	}
	return buf.Bytes(), nil
}
