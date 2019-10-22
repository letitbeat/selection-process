package tag

type TaskTags struct {
	TaskID string   `json:"id" bson:"_id"`
	Tags   []string `json:"tags" bson:"tags"`
}
