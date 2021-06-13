package store_test

import (
	"testing"

	"github.com/furrygem/authentication_server/internal/app/model"
	"github.com/furrygem/authentication_server/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func Test_DBOpen(t *testing.T) {
	sc := store.NewConfig()
	sc.DbPasswordFile = "testing_password.txt"
	sc.DbPort = 15432
	sc.DbAddr = "127.0.0.1"
	sc.DbDB = "tests"
	sc.SSLMode = "disable"
	st, err := store.Open(sc)
	assert.NoError(t, err, "Opening database")
	assert.NotNil(t, st, "Pointer to store instance is not nil")
}

func Test_QueryStatementFromMap(t *testing.T) {
	m := model.NewUser()
	m.ClearPassword = "secret"
	m.Username = "username"
	m.Prepare()
	sc := store.NewConfig()
	sc.DbPasswordFile = "testing_password.txt"
	sc.DbPort = 15432
	sc.DbAddr = "127.0.0.1"
	sc.DbDB = "tests"
	sc.SSLMode = "disable"
	st, err := store.Open(sc)
	assert.NoError(t, err, "Opening database")
	assert.NotNil(t, st, "Pointer to store instance is not nil")
	map1 := m.MapFromModel()
	stmt, err := st.QueryStatementFromMap(map1)
	assert.NoError(t, err, "Creating statement")
	assert.NotNil(t, stmt, "stmt object must not be nil")
}
