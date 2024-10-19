package data

import "os"

const DB_HOST = "db"

var DB_USER = os.Getenv("POSTGRES_USER")
var DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
var DB_NAME = os.Getenv("POSTGRES_DB")

const DB_CONTEXT string = "db"