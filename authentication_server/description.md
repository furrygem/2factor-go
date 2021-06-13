# Authentication API service
Here is the logic that authenticates and registers clients


## Configuration file
Default path for configuration file is authentication_server/configs/config.toml

Server is setting default values to all varibles that have not been set in config file. If you need to change something just add it to config file

### Configuration parameters

- bind_addr *string* - address to bind authentication server to. Must include host and port to bind to and be in url format "host:port". By default is set to "127.0.0.1"
- log_level *string* - log level for logger. Must contain interpretable by logger log level ("Trace" | "Debug", "Info", "Warning", "Error", "Fatal" | "Panic"). See official sirupsen logrus documentation at [github readme file for it]: https://github.com/sirupsen/logrus/blob/master/README.md. By default is set to "info".
**\[Store Config\]** 
- db_protocol *string* - protocol to use to communicate with database. By default is set to "postgres". **Supported drivers**: "postgres". You can add driver you need manually if it is not in the list, but server may not be stable.
- db_addr *string* - address of the database server to connect to. By default is set to "db". If running in docker-compose db is the DNS name for default database.
- db_port *int* - port of the database server is serving on. By default is set to default postgres port (5432)
- db_user *string* - username for provided db. By default is set to default postgres user ("postgres")
- db_password_file *string* - full path to the file containing password to database server. For security reasons recommended passing it as docker secret. By default is set to /run/secrets/dbpassword.txt.
- db_database *string* - path to the database on the database server. By default is set to "2factor". 
- db_sslmode *string* - sslmode for connecting to database. must be (disable|allow|prefer|require|verify-ca|verify-full) [see availiable sslmodes by postgres](https://www.postgresql.org/docs/11/libpq-ssl.html#LIBPQ-SSL-PROTECTION)

## Running tests
To run tests you should execute these commands:
```bash
make tests_start_docker
make tests
make clean_tests
```