package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "modernc.org/sqlite"

	parcelhttp "github.com/Yandex-Practicum/42-docker-final/internal/handler/http"
	"github.com/Yandex-Practicum/42-docker-final/internal/repository/sqlite"
	"github.com/Yandex-Practicum/42-docker-final/internal/service"
)

func main() {
	db, err := sql.Open("sqlite", "tracker.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	store := sqlite.NewParcelStore(db)
	svc := service.NewParcelService(store)

	mux := http.NewServeMux()
	handler := parcelhttp.NewParcelHandler(svc)
	handler.Register(mux)

	fmt.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}
}
