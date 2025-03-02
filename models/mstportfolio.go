package models

import (
	"time"
)

type MstPortfolio struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Description string `json:"description"`
	BackendStack string `json:"backend_stack"`
	FrontendStack string `json:"frontend_stack"`
	DatabaseStack string `json:"database_stack"`
	DeploymentStack string `json:"deployment_stack"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}