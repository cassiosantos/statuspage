package models

import "github.com/globalsign/mgo/bson"

// Client is a structure that identify a group of resources by a name
type Client struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string        `json:"name"`
	Resources []Component   `json:"resources"`
}
