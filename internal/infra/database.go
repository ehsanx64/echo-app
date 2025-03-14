package infra

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"

	"echo-app/internal/user/repository"
)

func DatabaseLoad() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=localhost user=echoapp password=echoapp dbname=echoapp")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := repository.New(conn)

	// list all users
	users, err := queries.Fetch(ctx)
	if err != nil {
		return err
	}
	log.Println(users)

	// create a user
	insertedUser, err := queries.Create(ctx, "User 1")
	if err != nil {
		return err
	}
	log.Println(insertedUser)

	// get the user we just inserted
	theUser, err := queries.Get(ctx, 2)
	if err != nil {
		return err
	}
	fmt.Println(theUser.Name)

	// prints true
	// log.Println(reflect.DeepEqual(insertedUser, theUser))
	return nil
}
