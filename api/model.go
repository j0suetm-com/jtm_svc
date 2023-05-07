package api

import "go.mongodb.org/mongo-driver/bson/primitive"

type LiteArtifact struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Title        string             `json:"title" bson:"title"`
	ExternalLink string             `json:"external_link" bson:"external_link"`
	Tags         []string           `json:"tags" bson:"tags"`
	HeaderId     primitive.ObjectID `json:"header_id" bson:"header_id"`
}

type Artifact struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Title        string             `json:"title" bson:"title"`
	ExternalLink string             `json:"external_link" bson:"external_link"`
	Tags         []string           `json:"tags" bson:"tags"`
	HeaderId     primitive.ObjectID `json:"header_id" bson:"header_id"`
	Content      []byte             `json:"content" bson:"content"`
}
