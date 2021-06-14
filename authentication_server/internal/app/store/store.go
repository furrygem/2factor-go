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

// Entrypoint to sore module. It takes store config as an arugument and returns pointer to configured store instance.
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

	stmt, err := s.QueryStatementFromMap("SELECT * FROM", model_map, "")
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
		if err := rows.Scan(&mdl.ID,
			&mdl.Username,
			&mdl.DiscordUsername,
			&mdl.FirstName,
			&mdl.Password,
		); err != nil {
			return nil, err
		}
		// adding pointer to model to result
		result = append(result, m)
	}

	return result, nil
}

// Inserts new user with specified values to users table and populates input model with values returnes form query.
// Can return error
func (s *Store) CreateUserWithModel(mdl *model.User) error {
	row := s.db.QueryRow("INSERT INTO users(username, discord_username, first_name, hash_password) VALUES ($1, $2, $3, $4) RETURNING *", mdl.Username, mdl.DiscordUsername, mdl.FirstName, mdl.Password)
	if err := row.Scan(&mdl.ID,
		&mdl.Username,
		&mdl.DiscordUsername,
		&mdl.FirstName,
		&mdl.Password,
	); err != nil {
		return err
	}
	return nil
}

func (s *Store) UpdateUserWithModel(target_model *model.User, updated_model *model.User) error {
	users_count, err := s.GetCountOfUsersWithModel(target_model)
	if err != nil {
		return err
	}
	if users_count > 1 {
		return errMultipleUsersFetchedByTargetInUpdateQuery
	}
	// diff_map := s.ModelDiff(target_model.MapFromModel(), updated_model.MapFromModel())
	stmt, err := s.UpdateStatementFromMap(updated_model.MapFromModel(), target_model.MapFromModel(), "RETURNING *")
	if err != nil {
		return err
	}
	var args []interface{}
	for _, val := range updated_model.MapFromModel() {
		args = append(args, val)
	}
	for _, val := range target_model.MapFromModel() {
		args = append(args, val)
	}
	row := stmt.QueryRow(args...)
	err = row.Scan(&updated_model.ID,
		&updated_model.Username,
		&updated_model.DiscordUsername,
		&updated_model.FirstName,
		&updated_model.Password,
	)
	if err != nil {
		return err
	}
	return nil
}

// Deletes entry from table by quering it with provided model. Input data must point only to 1 row.
// Can return error and populates input model with returned data from query
func (s *Store) DeleteUserWithModel(mdl *model.User) error {
	users_count, err := s.GetCountOfUsersWithModel(mdl)
	if err != nil {
		return err
	}
	if users_count > 1 {
		return errMultipleUsersFetchedByTargetInDeleteQuery
	}
	model_map := mdl.MapFromModel()
	stmt, err := s.QueryStatementFromMap("DELETE", model_map, "RETURNING *")
	if err != nil {
		return err
	}
	var query_args []interface{}
	for _, val := range model_map {
		query_args = append(query_args, val)
	}
	row := stmt.QueryRow(query_args...)
	err = row.Scan(
		&mdl.ID,
		&mdl.Username,
		&mdl.DiscordUsername,
		&mdl.FirstName,
		&mdl.Password,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetCountOfUsersWithModel(mdl *model.User) (int, error) {
	var (
		count      int
		query_args []interface{}
	)
	model_map := mdl.MapFromModel()
	stmt, err := s.QueryStatementFromMap("SELECT COUNT(*) FROM", model_map, "")
	if err != nil {
		return -1, err
	}
	for _, val := range model_map {
		query_args = append(query_args, val)
	}
	row := stmt.QueryRow(query_args...)
	err = row.Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

// Helper function to validate and set sslmode.
func (s *Store) issslmode(sslmode string, sslmodes_list []string) bool {
	return helpers.StringInSlice(sslmode, sslmodes_list)
}
