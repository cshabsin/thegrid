package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// Client is a client for interacting with Firebase Firestore.

type Client struct {
	*firestore.Client
}

// NewClient creates a new Firestore client.
func NewClient(ctx context.Context, projectID string, opts ...option.ClientOption) (*Client, error) {
	client, err := firestore.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}
