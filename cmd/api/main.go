package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tracker-tv/actor-api/internal/data"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/actors", func(w http.ResponseWriter, r *http.Request) {
		pgxTs := pgtype.Timestamp{Time: time.Now().UTC(), Valid: true}
		formattedStr := pgxTs.Time.Format("2006-01-02:15:04:05")
		parsedTime, _ := time.Parse("2006-01-02:15:04:05", formattedStr)

		actors := []data.Actor{
			{ID: 1, Name: "John Doe", CreatedAt: pgtype.Timestamp{Time: parsedTime, Valid: true}},
			{ID: 2, Name: "Jane Doe", CreatedAt: pgtype.Timestamp{Time: parsedTime, Valid: true}},
		}

		js, err := json.Marshal(actors)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	log.Fatalln(http.ListenAndServe(":8080", mux))
}
