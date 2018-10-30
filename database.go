package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func init() {

	fmt.Println("INIT DB!!!")

	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/betsmart?sslmode=disable")
	if err != nil {
		log.Fatal("Invalid DB config:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("DB unreachable:", err)
	}
}

func insertGame(game Game) {

	fmt.Println("Trying to insert: " + game.Id + " " + game.Sport + " " + game.Home + " " + game.Away + " " + game.Ts.String())

	sqlStatement := `
INSERT INTO public.games(id, sport, home, away, ts)
VALUES ($1, $2, $3, $4, $5)
RETURNING id`
	id := ""
	err := db.QueryRow(sqlStatement, game.Id, game.Sport, game.Home, game.Away, game.Ts).Scan(&id)

	if err != nil {
		//panic(err)
		fmt.Println("AN ERROR OCCORED", err)
	}
	fmt.Println("New game ID is:", id)
}

func selectGames() []Game {

	fmt.Println("Trying to get all * Games")

	rows, err := db.Query("SELECT id, sport, home, away, ts FROM games")

	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()

	var games []Game

	for rows.Next() {

		game := Game{}

		err = rows.Scan(&game.Id, &game.Sport, &game.Home, &game.Away, &game.Ts)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(game.Id)
		games = append(games, game)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return games
}
