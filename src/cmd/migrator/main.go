package main

import (
	"context"
	"os"
	"tenders-management/migrator"
)

func main() {
	migrator.Migrate(context.Background(), os.Stdout, os.Args[1:], os.LookupEnv, "postgres")
}
