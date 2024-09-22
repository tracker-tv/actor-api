package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tracker-tv/actor-api/internal/data"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/actors", func(w http.ResponseWriter, r *http.Request) {
		pgxTs := pgtype.Timestamp{Time: time.Now().UTC(), Valid: true}
		formattedStr := pgxTs.Time.Format("2006-01-02:15:04:05")
		parsedTime, _ := time.Parse("2006-01-02:15:04:05", formattedStr)

		actors := []data.Actor{
			{ID: 1, Name: "John Doe", CreatedAt: pgtype.Timestamp{Time: parsedTime, Valid: true}},
			{ID: 2, Name: "Jane Doe", CreatedAt: pgtype.Timestamp{Time: parsedTime, Valid: true}},
		}

		err := app.writeJSON(w, http.StatusOK, actors, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	})

	logger.Info("starting server", "port", 8080)

	log.Fatalln(http.ListenAndServe(":8080", mux))
}
