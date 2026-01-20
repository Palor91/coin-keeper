package handlers

import (
	"coin-keeper/internal/server/handlers/requirements"
)

type HandlersKeeper struct {
	dbEngine requirements.RequiredDatabase
}

func NewHandlersKeeper(dbEngine requirements.RequiredDatabase) *HandlersKeeper {
	return &HandlersKeeper{
		dbEngine: dbEngine,
	}
}
