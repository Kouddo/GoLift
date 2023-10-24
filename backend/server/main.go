package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

func main() {

	//TODO make these environment variables, dont store in code DONOT PUSH THESE
	connecString := "eat poo"
	//key := "poo poo goo goo"

	conn, err := pgx.Connect(context.Background(), connecString)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	containsBool, err := checkUserNameExists(conn, "abhi")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(containsBool)

	//router := gin.Default()
	//
	//router.Use(gin.Logger())
	//router.Use(gin.Recovery())
	//
	//router.GET("/default", getDefaultResponse)
	//
	//router.Run("0.0.0.0:8080")

}
