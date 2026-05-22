package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "modernc.org/sqlite"

	adapterhttp "github.com/Yandex-Practicum/42-docker-final/internal/adapter/http"
	"github.com/Yandex-Practicum/42-docker-final/internal/domain/usecase"
	"github.com/Yandex-Practicum/42-docker-final/internal/infrastructure/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "tracker.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() { _ = db.Close() }()

	store := sqlite.NewParcelStore(db)
	svc := usecase.NewParcelService(store)

	mux := http.NewServeMux()
	handler := adapterhttp.NewParcelHandler(svc)
	handler.Register(mux)

	fmt.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}
}
