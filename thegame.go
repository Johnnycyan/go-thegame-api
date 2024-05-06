package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	port := ":8027"
	sub := "/"
	handler := func(w http.ResponseWriter, r *http.Request) {
		output := getOutput(r)
		fmt.Fprintf(w, output)
	}
	fmt.Println("Listening on http://localhost" + port + sub)
	http.HandleFunc(sub, handler)
	http.ListenAndServe(port, nil)
}

func getOutput(r *http.Request) string {
	chanceArg := r.URL.Query().Get("chance")
	if chanceArg == "" {
		return "Please provide a chance argument"
	}
	chanceFloat, _ := strconv.ParseFloat(chanceArg, 64)
	win := r.URL.Query().Get("win")
	if win == "" {
		return "Please provide a win emote"
	}
	lose := r.URL.Query().Get("lose")
	if lose == "" {
		return "Please provide a lose emote"
	}
	randomNumber := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumberInt := randomNumber.Intn(1000)
	randomNumberFloat := float64(randomNumberInt) / 10
	doWin := randomNumberFloat <= chanceFloat
	if doWin {
		return "You just won The Game " + win
	}
	return "You just lost The Game " + lose
}
