package models

type User struct {
	Email          string `json:"id"`
	HashedPassword []byte `json:"hashedPassword"`
	Salt           []byte `json:"salt"`
}
