package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("DEV_DB_USERNAME"),
		os.Getenv("DEV_DB_PASSWORD"),
		os.Getenv("DEV_DB_NAME"))

	a.Run(":8080")
}
