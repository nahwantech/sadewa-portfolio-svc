package graph

import (
	"context"
	"fmt"

	"sadewa-portfolio-svc/config" // ✅ Correct way to import database
	"sadewa-portfolio-svc/graph/model"
)

// Fetch all mst_portfolios
func (r *Resolver) MstPortfolios(ctx context.Context) ([]*model.MstPortfolio, error) {
	var mstportfolios []model.MstPortfolio

	// Use `config.DB` for database queries
	if err := config.DB.Find(&mstportfolios).Error; err != nil {
		return nil, err
	}

	// Convert database models to GraphQL models
	var result []*model.MstPortfolio
	for _, mp := range mstportfolios {
		result = append(result, &model.MstPortfolio{
			ID:              fmt.Sprintf("%d", mp.ID),
			Title:           mp.Title,
			Description:     mp.Description,
			BackendStack:    mp.BackendStack,
			FrontendStack:   mp.FrontendStack,
			DatabaseStack:   mp.DatabaseStack,
			DeploymentStack: mp.DeploymentStack,
			CreatedAt:       mp.CreatedAt,
			CreatedBy:       mp.CreatedBy,
			UpdatedAt:       mp.UpdatedAt,
			UpdatedBy:       mp.UpdatedBy,
		})
	}

	return result, nil
}
