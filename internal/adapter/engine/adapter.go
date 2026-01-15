package engine

import "coin-keeper/internal/adapter/requirements"

type Adapter struct {
	dbEngine requirements.RequiredDatabase
}

func NewAdapter(dbEngine requirements.RequiredDatabase) (e *Adapter) {
	return &Adapter{dbEngine: dbEngine}
}
