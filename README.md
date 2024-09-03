Go simple crud & REST API service. Golang, PostgreSQL, GoMock.

## Installation
```bash
git clone https://github.com/hectorbarbosa/filmoteka.git
```
```bash
cd filmoteka
```
## Configuration
Adjust postgres credentials in `Makefile` (migrateup & migratedown) and in the config file (/config/apiserver.toml)
## Create database and tables
```bash
make createdb
```
```bash
make migrateup
```
## Running the service
```bash
make
```
```bash
make run
```
## Running tests 
```bash
make test
```

