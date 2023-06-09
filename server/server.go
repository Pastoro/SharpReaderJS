package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)
var dbpool, err = pgxpool.New(context.Background(),`user=postgres password=Pastori07@ host=localhost port=5432 dbname=Dictionary`)
func main(){
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var word string
	var definition string
	fmt.Println("Type a query:")
	fmt.Scanln(&word)
	query := fmt.Sprintf("select parola, definizione from it where parola='%s'",word)
	err = dbpool.QueryRow(context.Background(), query).Scan(&word, &definition)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(word, definition)
	http.HandleFunc("/", fetchWord)
	http.ListenAndServe(":8030", nil)



}

type user struct {

}

type wordDataGet struct {
	WordId string `json:"wordId"`
	CardStatus int `json:"cardStatus"`
	Learned  bool `json:"learned"`
	CardId string `json:"cardId"`
}

type wordDataSend struct {
	WordId string `json:"wordId"`
	CardStatus int `json:"cardStatus"`
	Learned  bool `json:"learned"`
	CardId string `json:"cardId"`
	Term string `json:"term"`
	Url string `json:"url"`
	TranslatedTerm string `json:"translatedTerm"`
}

func fetchWord(w http.ResponseWriter, request *http.Request) {


	//ParseForm is required to get the query parameters from the GET request
	request.ParseForm()
	fmt.Println(request.Form.Get("QueryWord"))
	fmt.Println(request.Form["QueryWord"])
	var word = request.Form.Get("QueryWord")
	var definition string

	query := fmt.Sprintf("select parola, definizione from it where parola='%s'",word)
	err = dbpool.QueryRow(context.Background(), query).Scan(&word, &definition)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(word, definition)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(definition); err != nil {
        panic(err)
    }
}