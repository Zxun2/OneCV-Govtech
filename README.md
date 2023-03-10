# OneCV Govtech Technical Assessment

By: Lee Zong Xun

## Introduction

This is a REST API built using Go and MySQL. It is a simple application that allows teachers to register students and retrieve a list of students common to a given list of teachers.

## Set up guide

This project is built using docker. To run the project, you will need to have docker installed on your machine.

1. Clone the repository
2. Install the dependencies
3. Run `docker pull mysql:8.0` to pull the MySQL docker image
4. Run `make setup` to run the MySQL database
   (**Note: Please give it some time to spin up the database before running the next command!**)
5. Run `make createdb` to create the database
6. Finally, run `make run` to run the application

The needed environment variables are given in the `app.env` file. You can change the values of the environment variables in the `app.env` file to suit your needs.

## Testing guide

To run the tests, run `make test` in the root directory of the project. Ensure that the MySQL **testing** database is running before running the tests. You can run the tests using the following command:

```bash
make createtestdb # create the testing database
make test
```

## API Documentation

There are 4 main routes in the application:

1. `api/register` - Registers one or more students to a specified teacher.
2. `api/commonstudents` - Retrieves a list of students common to a given list of teachers.
3. `api/suspend` - Suspends a specified student.
4. `api/retrievefornotifications` - Retrieves a list of students who can receive a given notification.

## Folder Structure

```
.
├── Makefile
├── README.md
├── api
│   ├── common.go
│   ├── server.go
│   ├── student.go
│   └── teacher.go
├── app.env
├── db
│   ├── db.go
│   └── migrations
│       └── ...
├── errors
│   ├── badrequest.go
│   └── recordnotfound.go
├── main.go
├── models
│   ├── student.go
│   └── teacher.go
├── seed
│   └── seed.go
├── services
│   ├── student.go
│   └── teacher.go
├── sqlc.yaml
├── tests
│   ├── main_test.go
│   ├── student_test.go
│   └── teacher_test.go
└── utils
    ├── config.go
    ├── random.go
    └── utils.go
```

## Common Troubleshooting

You can stop the container and start it again using the following commands:

```bash
docker stop mysql-root # stop the container
make setup
```
