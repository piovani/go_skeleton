package api

import (
	"fmt"

	"github.com/piovani/go_skeleton/infra/database"
)

type Api struct {
	db database.DatabaseContract
}

func NewApi() *Api {
	return &Api{
		db: database.NewDatabase(),
	}
}

func (a *Api) Start() error {
	fmt.Println("AQUI")
	return nil
}
