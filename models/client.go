package models

// Client is a structure that identify a group of resources by a name
type Client struct {
	Ref       string   `json:"ref,omitempty"`
	Name      string   `json:"name"`
	Resources []string `json:"resource_refs"`
}
