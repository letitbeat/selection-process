package user

import "time"

// User represents a "Sider" within the system, is used to transfer
// data between DB (MongoDB) and services.
type User struct {
	ID               string     `json:"siderId" bson:"_id"`
	CreatedAt        *time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt        *time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	FirstName        string     `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName         string     `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Tags             []string   `json:"tags,omitempty" bson:"tags,omitempty"`
	TaskApplications []string   `json:"taskApplications,omitempty" bson:"taskApplications,omitempty"`
	Score            float64    `json:"score"`
}
