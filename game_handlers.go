package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"
)

//Game Struct
type Game struct {
	Id    string    `json:"id"`
	Sport string    `json:"sport"`
	Home  string    `json:"home"`
	Away  string    `json:"away"`
	Ts    time.Time `json:"ts"`
}

func generateUUID() string {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
	return string(out)
}

func getGameHandler(w http.ResponseWriter, r *http.Request) {

	//connectToDB
	//setupDB()

	//Convert the "birds" variable to json
	var games = selectGames()

	fmt.Println(fmt.Errorf("RETURNING GAME: %s", games[0].Home))

	gamesBytes, err := json.Marshal(games)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("BIRDS GET: " + string(gamesBytes))
	// If all goes well, write the JSON list of birds to the response
	w.Write(gamesBytes)
}

func createGameHandler(w http.ResponseWriter, r *http.Request) {

	//setupDB()

	// Create a new instance of Bird
	game := Game{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the bird from the form info

	game.Id = generateUUID()
	game.Sport = "BASKETBALL"
	game.Home = r.Form.Get("home")
	game.Away = r.Form.Get("away")
	game.Ts = time.Now()

	// Append our existing list of birds with a new entry
	//games = append(games, game)
	insertGame(game)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/", http.StatusFound)
}
