package task

import "time"

type Task struct {
	ID          string      `json:"id" bson:"_id"`
	Name        string      `json:"name" bson:"name"`
	Description string      `json:"description" bson:"description"`
	CreatedAt   *time.Time  `json:"createdAt" bson:"createdAt"`
	UpdateAt    *time.Time  `json:"updatedAt" bson:"updatedAt"`
	Country     string      `json:"country" bson:"country"`
	Tags        []string    `json:"tags" bson:"tags"`
	Applicants  []Applicant `json:"applicants" bson:"applicants"`
}

type Applicant struct {
	ID string `json:"applicantId" bson:"applicantId"`
}
