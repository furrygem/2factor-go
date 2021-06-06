package store

type Config struct {
	DbProtocol     string `toml:"db_protocol"`      // ex. db_protocol = "postgres"
	DbAddr         string `toml:"db_addr"`          // ex. db_addr = "db"
	DbPort         int    `toml:"db_port"`          //  ex. db_port = 5432
	DbUser         string `toml:"db_user"`          // ex. db_user = "postgres"
	DbPasswordFile string `toml:"db_password_file"` /* db_password_file - name of the file containing password for database. Recommended passing password as
	Docker secret | ex. db_password_file = "/run/secrets/dbpassword.txt" */
	DbDB string `toml:"db_database"` // ex. db_database = "2factor"
	/*
		The example strings above will be formatted as "postgres://postgres:(password_file_content)@db:5433/2factor"
	*/
}

/*
	Returns pointer to a new config struct populated with default values
*/
func NewConfig() *Config {
	// returning addresss of Config struct populated with default values
	return &Config{
		DbProtocol:     "postgres",
		DbAddr:         "db",
		DbPort:         5432,
		DbUser:         "postgres",
		DbPasswordFile: "/run/secrets/dbpassword.txt",
		DbDB:           "2factor",
	}
}
