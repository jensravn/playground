package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

func deleteAll(projectID, kind, namespace string) error {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("new client: %w", err)
	}
	query := datastore.NewQuery(kind).Namespace(namespace).KeysOnly()
	keys, err := client.GetAll(ctx, query, nil)
	if err != nil {
		return fmt.Errorf("get all: %w", err)
	}
	err = client.DeleteMulti(ctx, keys)
	if err != nil {
		return fmt.Errorf("delete multi: %w", err)
	}
	return nil
}
