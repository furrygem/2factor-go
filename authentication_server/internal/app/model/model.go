// package model provides models for different db tables
package model

import (
	"reflect"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              int64  `json:"id" col:"id"`
	Username        string `json:"username" col:"username"`
	DiscordUsername string `json:"discord_username" col:"discord_username"`
	FirstName       string `json:"first_name" col:"first_name"`
	ClearPassword   string `json:"clear_password"`
	Password        string `json:"password" col:"hash_password"`
}

func NewUser() *User {
	return &User{}
}

// Method Map from model converts UserInstance to map[string]inteface{}. if Model value is equal to zero then it will not be represented in the result map
func (m *User) MapFromModel() map[string]interface{} {
	t := reflect.TypeOf(*m)
	v := reflect.ValueOf(*m)
	res := make(map[string]interface{}) // Creating result map that will be returned
	for i := 0; i < v.NumField(); i++ { // Iterating through every field of the target struct
		// if field value is zero continuing the loop
		if v.FieldByName(t.Field(i).Name).IsZero() {
			continue
		}
		// if not
		typ := v.FieldByName(t.Field(i).Name).Kind().String() // Getting type of the field and converting it to string
		tag := t.Field(i).Tag.Get("col")                      // getting "col" tag
		value := v.FieldByName(t.Field(i).Name)
		if typ == "string" {
			res[tag] = value.String()
		} else if typ == "int" {
			res[tag] = value.Int()
		} else {
			// setting res map with col tag value as the key to the value represented as interface{}
			res[tag] = value.Interface()
		}
	}
	// returning result map
	return res
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		validation.Field(&u.Username, validation.Required, validation.Length(1, 20)),
		validation.Field(&u.DiscordUsername, validation.Required, validation.Match(regexp.MustCompile("^.{3,32}#[0-9]{4}$"))),
		validation.Field(&u.FirstName, validation.Length(1, 25)),
		validation.Field(&u.ClearPassword, validation.Length(1, 64)),
	)
}

// Method Prepare hashes writes to Password field of the User model bcrypt hashed value if ClearPassword field.
// And then empties ClearPassword field by setting it to empty string.
// Returnes error.
func (u *User) Prepare() error {
	// Generating bcrypt hash from ClearPassword field and catching error
	password, err := bcrypt.GenerateFromPassword([]byte(u.ClearPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(password)
	u.ClearPassword = ""
	return nil
}

// If provided password is ok returning ok returning true, else returning false
func (u *User) CompareHashAndPassword(passwordToCheckAgainst string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passwordToCheckAgainst))
	return err == nil
}
