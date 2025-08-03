package dataservice

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/cshabsin/thegrid/apps/explorers/data"
	"google.golang.org/api/iterator"
)

// Client holds a connection to the Firestore database.
type Client struct {
	fs *firestore.Client
}

// NewClient creates a new client connected to Firestore.
// IMPORTANT: Credentials should be handled via environment variables
// (e.g., GOOGLE_APPLICATION_CREDENTIALS), not checked into code.
func NewClient(ctx context.Context) (*Client, error) {
	conf := &firebase.Config{ProjectID: "shabsin-thegrid"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating firestore client: %v", err)
	}
	return &Client{fs: client}, nil
}

// GetSystems retrieves all documents from the 'systems' collection.
func (c *Client) GetSystems(ctx context.Context) ([]*data.SystemData, error) {
	var systems []*data.SystemData
	iter := c.fs.Collection("systems").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var sys data.SystemData
		if err := doc.DataTo(&sys); err != nil {
			return nil, err
		}
		systems = append(systems, &sys)
	}
	return systems, nil
}

// GetPaths retrieves all documents from the 'paths' collection.
func (c *Client) GetPaths(ctx context.Context) ([]data.PathSegment, error) {
	var paths []data.PathSegment
	iter := c.fs.Collection("paths").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var path data.PathSegment
		if err := doc.DataTo(&path); err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}
	return paths, nil
}

func (c *Client) AddSystem(ctx context.Context, sys *data.SystemData) error {
	_, _, err := c.fs.Collection("systems").Add(ctx, sys)
	return err
}

func (c *Client) AddPath(ctx context.Context, path *data.PathSegment) error {
	_, _, err := c.fs.Collection("paths").Add(ctx, path)
	return err
}

func (c *Client) Close() {
	c.fs.Close()
}
