package user

import (
	"context"
	"log"

	"github.com/letitbeat/selection-process/pkg/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository which fetches User data from MongoDB.
type Repository struct {
	DBURI string
}

// GetByID returns a User object given its ID if matched,
// otherwise error.
func (r Repository) GetByID(ID string) (User, error) {
	ctx := context.Background()
	client, err := db.Client(ctx, r.DBURI)

	if err != nil {
		log.Printf("error connecting to DB %s", r.DBURI)
	}

	collection := client.Database("main").Collection("users")

	filter := bson.D{{"_id", ID}}

	var user User
	err = collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		log.Printf("error fetching user with ID: %s from DB", ID)
		return user, err
	}

	return user, nil
}

// GetByIDs returns a User object list given its IDs if matched,
// otherwise error.
func (r Repository) GetByIDs(IDs []string) ([]User, error) {
	ctx := context.Background()
	client, err := db.Client(ctx, r.DBURI)

	if err != nil {
		log.Printf("error connecting to DB %s", r.DBURI)
	}

	collection := client.Database("main").Collection("users")

	var users []User

	filter := bson.D{{"_id", bson.D{{"$in", IDs}}}}
	projection := bson.D{{"_id", 1}, {"firstName", 2}, {"lastName", 3}, {"tags", 4}}
	cursor, err := collection.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		log.Println("error getting collection", err.Error())
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		log.Println("error getting data from cursor", err.Error())
	}
	return users, nil
}
