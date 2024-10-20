// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Node interface {
	IsNode()
	GetID() string
}

// Post
type Post struct {
	// The ID of an object.
	ID string `json:"id"`
	// The image of Post.
	Image string `json:"image"`
	// The caption of Post.
	Caption *string `json:"caption,omitempty"`
	// The User who posted.
	PostedBy *User `json:"postedBy"`
}

func (Post) IsNode()            {}
func (this Post) GetID() string { return this.ID }

// Root Query.
type Query struct {
}

// User
type User struct {
	// The ID of an object.
	ID string `json:"id"`
	// The Name of User.
	Name string `json:"name"`
	// The Image of User.
	Image *string `json:"image,omitempty"`
}

func (User) IsNode()            {}
func (this User) GetID() string { return this.ID }
