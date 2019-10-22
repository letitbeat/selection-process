package tag

import (
	"context"
	"log"

	"github.com/letitbeat/selection-process/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository struct {
	DBURI string
}

func (r *Repository) GetByTaskID(taskID string) (TaskTags, error) {
	ctx := context.Background()
	client, err := db.Client(ctx, r.DBURI)

	if err != nil {
		log.Println("error connection to DB")
	}

	collection := client.Database("tags").Collection("tasksTags")

	filter := bson.D{{"_id", taskID}}

	var taskTags TaskTags
	err = collection.FindOne(ctx, filter).Decode(&taskTags)

	if err != nil {
		log.Printf("error fetching tags for task ID: %s from DB", taskID)
		return taskTags, err
	}

	return taskTags, nil
}
