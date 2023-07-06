package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(mongouri string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// TODO: change to env value
	clientOptions := options.Client().ApplyURI(mongouri)
	client, _ := mongo.Connect(ctx, clientOptions)
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database("albo"), nil
}
