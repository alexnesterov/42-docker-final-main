package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"

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

	client := 1
	address := "Псков, д. Пушкина, ул. Колотушкина, д. 5"
	p, err := svc.Register(client, address)
	if err != nil {
		fmt.Println(err)
		return
	}

	newAddress := "Саратов, д. Верхние Зори, ул. Козлова, д. 25"
	err = svc.ChangeAddress(p.Number, newAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.NextStatus(p.Number)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.PrintClientParcels(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Delete(p.Number)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.PrintClientParcels(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	p, err = svc.Register(client, address)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Delete(p.Number)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.PrintClientParcels(client)
	if err != nil {
		fmt.Println(err)
		return
	}
}
