package resolver

import (
	"sadewa-portfolio-svc/db"
)

// QueryResolver struct for handling queries
type QueryResolver struct {
	DB *db.Database
}

// Function to return DB connection
func (r *QueryResolver) DB() *db.Database {
	return r.DB
}
