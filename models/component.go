package models

import "github.com/globalsign/mgo/bson"

// Component represents a host, database or any resource with an address
type Component struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name" binding:"required"`
	Groups    []string      `json:"groups"`
	Address   string        `json:"address"`
	Incidents []Incident    `json:"incidents_history"`
}
