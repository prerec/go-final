# go-final

### How to configure the application:

The config directory contains the config.uml file. In it you can configure the port on which the server will be 
launched. The remaining parameters are provided in case the application migrates to a new database, and 
it is recommended not to change them.

`./configs/config.yml:`

```yml 
port: "7540"

### do not change that block
db:
db_driver: "sqlite3"
db_port: ""
db_username: ""
db_password: ""
db_name: "scheduler.db"
db_ssl_mode: ""
```

### How to start project:

```bash
cd go-final
```

- ### In a docker container:

```bash
make docker_build
```
```bash
make docker_run
```

- ### Locally:

Install the goose migrations utility if not already installed:

```bash
make goose
```

Launch the application:

```bash
make run
```

### How to test an application:
To configure tests, change the file: `./test/settings.go`

```go
var Port = 7540 // set the port according to the setting from the configuration file ./configs/config.yml
var DBFile = "../scheduler.db" // set the path to the database file in the project root
var FullNextDate = false // do not change since setting monthly repetitions is not implemented
var Search = true // set to true to test the search implementation
var Token = `` // do not change as it is not implemented
```

How to start tests:
```bash
go test ./tests
```