# Authentication API service
Here is the logic that authenticates and registers clients


## Configuration file
Default path for configuration file is authentication_server/configs/config.toml

Server is setting default values to all varibles that have not been set in config file. If you need to change something just add it to config file

### Configuration parameters

- bind_addr *string* - address to bind authentication server to. Must include host and port to bind to and be in url format "host:port". By default is set to "127.0.0.1"
- log_level *string* - log level for logger. Must contain interpretable by logger log level ("Trace" | "Debug", "Info", "Warning", "Error", "Fatal" | "Panic"). See official sirupsen logrus documentation at [github readme file for it](https://github.com/sirupsen/logrus/blob/master/README.md). By default is set to "info".

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
### API functionality
#### CRUD
authentication server API provides CRUD to access different tables of connected DB
##### Create
POST to /api/users JSON struct of new user to create it.\
**Request:**
```json
{
    "username": "required_useraname",
    "discord_username": "discord_username with regex ^.{3,32}#[0-9]{4}$ (default discord username syntax)",
    "first_name": "first name (not required to create user)" ,
    "clear_password": "password in clear. it will be hashed with bcrypt and added to db in hash form. clear password will not be stored",
}
```
**Response:**

Response is JSON object describing new user.
```json
{
    "id": "user id. integer. serial number. set by db",
    "username": "unique not null string",
    "discord_username": "unique not null string. discord username",
    "first_name": "user's first name. can be empty",
    "hash_password": "bcrypt hash of user password"
}
```

##### Read
GET request to /api/users specifying query data will return all users that suite in query arguments.\
For example:

**Request:**
```json
{
    "first_name": "John"
}
```
**Response:**

Response will be json list of JSON objects of all users suitable for query.
```json
[
    {
        "id": 123,
        "username": "username123",
        "discord_username": "unique not null string. discord username",
        "first_name": "John",
        "hash_password": "bcrypt hash of user password"
    }
    {
        "id": 223,
        "username": "username223",
        "discord_username": "unique not null string. discord username",
        "first_name": "John",
        "hash_password": "bcrypt hash of user password"
    }
]
```
##### Update
To update row you need to make PUT request to /api/users specifying query for target row (there must be only 1 user that suits provided query arguments) and specifying new values for each field that needed to be changed.

**Request:**
```json
{
    "target_user": {
        "id": 123,
        "first_name": "John"
    },
    "updated_values": {
        "first_name": "Eugene"
    }
}
```
**Response:**

This will return JSON object describing user with updated values.
```json
{
    "id": 123,
    "username": "username123",
    "discord_username": "user#1234",
    "first_name": "Eugene",
    "hash_password": "longlonghash"
}
```
##### Delete
To update row you need to make PUT request to /api/users specifying query for target row (there must be only 1 user that suits provided query arguments).

**Request:**
```json
{
    "id": 123,
    "first_name": "Eugene",
}
```
**Response:**

Response will contain JSON object describing user that just was deleted.
```json
{
    "id": 123,
    "username": "username123",
    "discord_username": "user#1234",
    "first_name": "Eugene",
    "hash_password": "longlonghash"
}
```