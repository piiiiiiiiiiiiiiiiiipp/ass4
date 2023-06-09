package main

import (
	"fmt"

	"awesomeProject3/pkg/store/postgres"
)

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()

	fmt.Println(conn.Pool.Stat())

	fmt.Println("Hello World!")
}
