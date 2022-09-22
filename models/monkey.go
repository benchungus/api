package models

//struct for Monkey object
type Monkey struct {
	Name  string `bson:"name"`
	Type  string `bson:"type"`
	Hobby string `bson:"hobby"`
}
