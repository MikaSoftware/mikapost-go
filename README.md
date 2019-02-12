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


## Documentation

* [Install on MacOS / Linux for development](/docs/Dev-1-Setup-and-Install.md)

* [(DevOps Part 1) Install on CentOS 7 for production ](/docs/DevOps-1-Setup-DigitalOcean-CentOS-7.md)

* [(DevOps Part 2) Setup project](/docs/DevOps-2-Setup-Project.md)

* [(DevOps Part 3) Setup Lets Encrypt](/docs/DevOps-3-Setup-Lets-Encrypt.md)

* [API Reference](/docs/Dev-x-API-Reference.md)
