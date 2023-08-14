package models

type ToDo struct {
	ID       string `json:"-" bson:"_id,omitempty"`
	Title    string `json:"title" bson:"title" validate:"required"`
	ActiveAt string `json:"activeAt" bson:"activeAt" validate:"required"`
	Status   string `json:"-" bson:"status"`
}
