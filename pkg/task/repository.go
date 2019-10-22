package task

import (
	"context"
	"log"

	"github.com/letitbeat/selection-process/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository struct {
	DBURI string
}

func (r *Repository) GetByID(ID string) (Task, error) {
	ctx := context.Background()
	client, err := db.Client(ctx, r.DBURI)

	if err != nil {
		log.Printf("error connecting to DB %s", r.DBURI)
	}

	collection := client.Database("main").Collection("tasks")

	filter := bson.D{{"_id", ID}}

	var task Task
	err = collection.FindOne(ctx, filter).Decode(&task)

	if err != nil {
		log.Printf("error fetching task ID: %s from DB", ID)
		return task, err
	}

	return task, nil
}
