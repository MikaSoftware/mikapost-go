# MikaPost (Golang)
**(Currently under development)**

MikaPost is time-series data *storage* and *sharing* platform. Create **”boxes”** where you can place **“things”** into to be populated with with any sort of **time-series data**.

What is a **box**? A box could be anything that you’d like to organization by. Here are a few examples:

* Census for the current year
* Financial record for a customer
* A particular science experiment
* Engineering data for a prototype
* Medical records for a patient
* Home automation for particular IoT devices
* Industrial site operating sensors

What is a **thing**? A thing could be anything you’d like to tag for the *time-series data* that you’ll be uploading. Here are a few examples:

* The hours spent by city council in meeting
* Customer purchase amount ($) over time
* The population size for a cell-culture over time
* Tensile strength of a material in an automotive product
* Apple HealthKit walking data
* Home temperature sensor
* Battery level of some machine

Current features include:

* Organize multiple things under a box.
* Grant read or read/write access to other users for a box / thing.
* Set a box or thing to be public to be used by the public or set them to be private to restrict access.
* Keep boxes and things private but share them with users who don't have an account through "private URLs".
* CRUD operations on time-series data
* User registration / Login

This repository is the go implementation of the web-service backend which is interacted with a JS single page app.

## Installation
1. Get the latest code.

  ```bash
  go get -u -v github.com/mikasoftware/mikapost-go
  ```


2. Open up ``postgre``s and run the following to setup the database for **development** or **production**.

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


3. Run the following to setup the database for **unit testing**.

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


4. Install the project dependencies.

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
