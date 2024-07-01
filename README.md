# Questions-go

## Installation

```bash
go get -u
go mod tidy
go mod download
```

## Overview
This is a server side aplication that provides a simple quizz with a short list of questions for the user to answer. 

The user is a single session user, which means that every time answers are submitted, the systems understands it as a new user. 

Furthermore, users can see their statistiscs. The system responds with three percents: 
    - Percent of users that got more correct answers. 
    - Percent of users that got the same number of correct answers. 
    - Percent of users that got fewer correct answers. 

You can talk with the API with a buit in CLI stored in /cli. For its usage you first need to run the project. 

```bash
cd app
go run .
```



