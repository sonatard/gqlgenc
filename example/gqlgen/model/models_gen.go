// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type Profile struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Query struct {
}

type User struct {
	ID      string   `json:"id"`
	Profile *Profile `json:"profile"`
}
