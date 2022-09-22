package models

type Monkey struct {
	//ID   primitive.ObjectID `bson:"_id"`
	Name  string `bson:"name"`
	Type  string `bson:"type"`
	Hobby string `bson:"hobby"`
}
