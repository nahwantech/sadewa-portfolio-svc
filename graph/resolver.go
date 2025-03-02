package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"sadewa-portfolio-svc/config"
	"gorm.io/gorm"
)

type Resolver struct{}

func (r *Resolver) DB() *gorm.DB {
	return database.DB
}
