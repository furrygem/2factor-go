package model_test

import (
	"fmt"
	"testing"

	"github.com/furrygem/authentication_server/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func Test_ModelPrepareMethod(t *testing.T) {
	m := model.NewUser()
	m.ClearPassword = "secret"
	m.Username = "username"
	m.Prepare()
	assert.Zero(t, m.ClearPassword, "Clear password should be equal to zero after calling Prepare method of model")
	x := m.CompareHashAndPassword("secret")
	fmt.Println(x)
	assert.True(t, m.CompareHashAndPassword("secret"), "Password validation check on valid password")
	assert.False(t, m.CompareHashAndPassword("secret1"), "Password validation check on invalid password")
}

func Test_MapFromModel(t *testing.T) {
	m := model.NewUser()
	m.ClearPassword = "secret"
	m.Username = "username"
	map1 := m.MapFromModel()
	assert.Equal(t, "username", map1["username"], "username in map must be valid")
	assert.Nil(t, map1["clear_password"], "clear_password must not be exported to map")
	assert.Nil(t, map1["hash_password"], "hash password must be nil before model.Prepare method call")
	m.Prepare()
	map2 := m.MapFromModel()
	assert.Equal(t, "username", map2["username"], "Username in map must be valid")
	assert.Nil(t, map2["clear_password"], "clear_password must not be exported to map")
	assert.Equal(t, m.Password, map2["hash_password"])
}
