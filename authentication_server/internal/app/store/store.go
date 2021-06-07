//Package store provides interface and logic for communicating with database
package store

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq" // importing postgres driver anonymously
)

// main struct for store package that will contain configured modules instances
type Store struct {
	db     *sql.DB
	Logger *logrus.Logger
}

// Entrypoint to sore module. It takes store config as an arugument and returns pointer to configured store instance
func Open(c *Config) (*Store, error) {
	var (
		db_address    string
		format_string string
		db_password   string
		store         *Store
	)
	store = &Store{}                      // creating empty store instance
	format_string = "%s://%s:%s@%s:%i/%s" // string containing format template for db url
	/* converting values from provided config to database address */

	/* Reading password file contnents and catching error */
	data, err := ioutil.ReadFile(c.DbPasswordFile)
	if err != nil {
		return nil, err //We do not returning any Store struct but returning error
	}
	db_password = strings.TrimSuffix(string(data), "\n") // setting content of the file to db_password varible and trimming newline in the suffix
	/*formatting address */
	db_address = fmt.Sprintf(format_string, c.DbProtocol, c.DbUser, db_password, c.DbAddr, c.DbPort, c.DbDB)
	/* openenig database in lazy mode and catching error  */

	db, err := sql.Open("postgres", db_address)
	if err != nil {
		return nil, err
	}
	/* pinging opened in lazy mode db to initialize it and catching error */
	db.Ping()
	if err != nil {
		return nil, err
	}
	store.db = db     // assigning db to store store instance
	return store, nil //returning pointer to initialized store instance and nil error

}
