package resolver

import (
	"context"
	"fmt"
	"sadewa-portfolio-svc/graph/model" // GraphQL model
	"sadewa-portfolio-svc/models"      // Database model
	"time"
)

// Fetch all mst_portfolios
func (r *QueryResolver) MstPortfolios(ctx context.Context) ([]*model.MstPortfolio, error) {
	var mstportfolios []models.MstPortfolio

	// Fetch from DB
	if err := r.DB().Find(&mstportfolios).Error; err != nil {
		return nil, err
	}

	// Convert database models to GraphQL models
	var result []*model.MstPortfolio
	for _, mp := range mstportfolios {
		converted := &model.MstPortfolio{
			ID:              fmt.Sprintf("%d", mp.ID), // Convert uint to string
			Title:           mp.Title,
			Description:     mp.Description,
			BackendStack:    mp.BackendStack,
			FrontendStack:   mp.FrontendStack,
			DatabaseStack:   mp.DatabaseStack,
			DeploymentStack: mp.DeploymentStack,
			CreatedAt:       mp.CreatedAt.Format(time.RFC3339), // Convert time.Time to string
			CreatedBy:       mp.CreatedBy,
			UpdatedAt:       mp.UpdatedAt.Format(time.RFC3339), // Convert time.Time to string
			UpdatedBy:       mp.UpdatedBy,
		}
		result = append(result, converted) // Correctly append GraphQL models
	}

	return result, nil
}
