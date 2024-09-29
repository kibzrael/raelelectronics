package data

import "os"

const DB_HOST = "db"

var DB_USER = os.Getenv("POSTGRES_USER")
var DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
var DB_NAME = os.Getenv("POSTGRES_DB")

type key string

const DB_CONTEXT key = "db"

const WG_CONTEXT key = "wg"
const WG_DETAIL_CONTEXT key = "wgDetail"

var API_KEY = os.Getenv("SEED_API_KEY")
