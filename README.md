# mikapost-go
The web platform for buying and selling comic books online. The ``go`` repository is the back-end service which powers a front-end service.

## Installation
1. Get the latest code.

  ```bash
  go get -u -v github.com/mikasoftware/mikapost-go
  ```


2. Open up postgres and run the following to setup the database for development or production.

  ```sql
  drop database mikapost_db;
  create database mikapost_db;
  \c mikapost_db;
  CREATE USER golang WITH PASSWORD 'YOUR_PASSWORD';
  GRANT ALL PRIVILEGES ON DATABASE mikapost_db to golang;
  ALTER USER golang CREATEDB;
  ALTER ROLE golang SUPERUSER;
  CREATE EXTENSION postgis;
  ```


3. Run the following to setup the database for unit testing.

  ```sql
  drop database mikapost_test_db;
  create database mikapost_test_db;
  \c mikapost_test_db;
  CREATE USER golang WITH PASSWORD 'YOUR_PASSWORD';
  GRANT ALL PRIVILEGES ON DATABASE mikapost_test_db to golang;
  ALTER USER golang CREATEDB;
  ALTER ROLE golang SUPERUSER;
  CREATE EXTENSION postgis;
  ```


4. Install our dependencies

  ```
  ./requirements.sh
  ```


5. Update environmental variables by running the following. Please change the variables to whatever you prefer.

  ```bash
  #!/bin/bash
  export MIKAPOST_GORM_CONFIG="postgres://golang:YOUR_PASSWORD@localhost/mikapost_db?sslmode=disable"
  export MIKAPOST_SECRET="YOUR_SECRET_RANDOM_STRING"
  export MIKAPOST_ADDRESS="127.0.0.1:8080"
  export TEST_MIKAPOST_GORM_CONFIG="postgres://golang:YOUR_PASSWORD@localhost/mikapost_test_db?sslmode=disable"
  ```
