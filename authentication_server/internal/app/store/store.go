//Package store provides interface and logic for communicating with database
package store

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/furrygem/authentication_server/internal/app/model"
	"github.com/sirupsen/logrus"

	"github.com/furrygem/authentication_server/internal/app/helpers"

	_ "github.com/lib/pq" // importing postgres driver anonymously
)

// main struct for store package that will contain configured modules instances
type Store struct {
	db     *sql.DB
	Logger *logrus.Logger
}

// Entrypoint to sore module. It takes store config as an arugument and returns pointer to configured store instance
func Open(c *Config) (*Store, error) {
	var AVAILABLE_SSLMODES = []string{
		"disable",
		"allow",
		"prefer",
		"require",
		"verify-ca",
		"verify-full",
	}

	var (
		db_address    string
		format_string string
		db_password   string
		store         *Store
	)
	store = &Store{}                                 // creating empty store instance
	format_string = "%s://%s:%s@%s:%d/%s?sslmode=%s" // string containing format template for db url
	/* converting values from provided config to database address */

	/* Reading password file contnents and catching error */
	data, err := ioutil.ReadFile(c.DbPasswordFile)
	if err != nil {
		return nil, err //We do not returning any Store struct but returning error
	}
	db_password = strings.TrimSuffix(string(data), "\n") // setting content of the file to db_password varible and trimming newline in the suffix

	if !store.issslmode(c.SSLMode, AVAILABLE_SSLMODES) {
		return nil, errNotSslMode
	}
	/*formatting address */
	db_address = fmt.Sprintf(format_string, c.DbProtocol, c.DbUser, db_password, c.DbAddr, c.DbPort, c.DbDB, c.SSLMode)
	/* openenig database in lazy mode and catching error  */

	db, err := sql.Open("postgres", db_address)
	if err != nil {
		return nil, err
	}
	/* pinging opened in lazy mode db to initialize it and catching error */
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	store.db = db     // assigning db to store store instance
	return store, nil //returning pointer to initialized store instance and nil error

}

func (s *Store) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUserByModel(mdl *model.User) ([]*model.User, error) {
	// getting user by data from provided model
	var (
		result     []*model.User
		query_args []interface{}
	)
	model_map := mdl.MapFromModel()

	stmt, err := s.QueryStatementFromMap(model_map)
	if err != nil {
		return nil, err
	}
	// quering multiple rows with stmt.Query
	for _, val := range model_map {
		query_args = append(query_args, val)
	}
	rows, err := stmt.Query(query_args...)
	if err != nil {
		return nil, err
	}
	// for every row
	for rows.Next() {
		// creating new user model
		m := model.NewUser()
		// scanning row data to it
		if err := rows.Scan(&m.ID, &m.Username, &m.Password); err != nil {
			return nil, err
		}
		// adding pointer to model to result
		result = append(result, m)
	}

	return result, nil
}

// Helper function to validate and set sslmode
func (s *Store) issslmode(sslmode string, sslmodes_list []string) bool {
	return helpers.StringInSlice(sslmode, sslmodes_list)
}
