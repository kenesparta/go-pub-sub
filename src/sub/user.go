package main

import "encoding/json"

// User is a struct representing newly registered users
type User struct {
	Username string
	Email    string
}

// MarshalBinary encodes the struct into a binary blob
func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

// UnmarshalBinary decodes the struct into a User
func (u *User) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, u); err != nil {
		return err
	}
	return nil
}

func (u *User) String() string {
	return "User: " + u.Username + " registered with Email: " + u.Email
}
