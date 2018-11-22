package models

import "github.com/globalsign/mgo/bson"

// Webhook is the definition of both a Incoming and Outgoing webhook
type Webhook struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"` //
	Type        string        `json:"type"`                    // Incoming || Outgoing types
	Enabled     bool          `json:"enabled"`                 //
	Token       string        `json:"token,omitempty"`         // Incoming: the endpoint Identification || Outgoing: Any
	Name        string        `json:"name"`                    // Shared
	Description string        `json:"description,omitempty"`   // Shared
	Target      string        `json:"url,omitempty"`           // Outgoing destination
	Trigger     string        `json:"trigger,omitempty"`       // Outgoing activation
	Payload     interface{}   `json:"data"`                    // Shared
}

const (
	// WebhookIncomingType identify webhook incoming type
	WebhookIncomingType = "incoming"
	// WebhookOutgoingType identify webhook outgoing type
	WebhookOutgoingType = "outgoing"
)
