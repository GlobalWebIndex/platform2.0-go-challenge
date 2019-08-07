package model

// User is the struct for the user model.
type User struct {
	ID     uint             `json:"id"`
	Assets map[string]Asset `json:"assets,omitempty"`
}
