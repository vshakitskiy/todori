package ent

import (
	ctx "context"
	f "fmt"
	"log"
	"todori/lib"

	_ "github.com/lib/pq"
)

var client *Client

func OpenClient() {
	go func() {
		
	}()
	host, port, user, password, dbname := 
		lib.Getenv("POSTGRES_HOST"), 
		lib.Getenv("POSTGRES_PORT"), 
		lib.Getenv("POSTGRES_USER"), 
		lib.Getenv("POSTGRES_PASSWORD"), 
		lib.Getenv("POSTGRES_DB")

	dsn := f.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		host, 
		port, 
		user, 
		password, 
		dbname,
	)

	client, err := Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed opening conn to postgres database: %v", err)
	}
	log.Println("Database connected")

	if err := client.Schema.Create(ctx.Background()); err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}
	log.Println(
		"Schema resources created", 
		"\n_____________________________________________",
	)
}

func GetClient() *Client {
	return client
}