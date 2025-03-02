package resolver

import(
	"fmt"

	"context"
	"sadewa-portfolio-svc/graph/model"
	"sadewa-portfolio-svc/model"
)

//Fetch all mst portfolio
func (r *queryResolver) MstPortfolio(ctx context.Context) ([]*model.MstPortfolio, error) {
	var mstportfolios []models.MstPortfolio
	if err != r.DB().Find(&mstportfolios).Error; err != nil {
		return nil, err
	}

	var result []*models.MstPortfolio
	for _, mp := range mstportfolios {
		result = append(result, &model.MstPortfolio{
			ID: fmt.Sprintf("%d", mp.ID),
			Title: mp.Title,
			Description: mp.Description,
			BackendStack: mp.BackendStack,
			FrontendStack: mp.FrontendStack,
			DatabaseStack: mp.DatabaseStack,
			DeploymentStack: mp.DeploymentStack,
			CreatedAt: mp.CreatedAt,
			CreatedBy: mp.CreatedBy,
			UpdatedAt: mp.UpdatedAt,
			UpdatedBy: mp.UpdatedBy,
		})
	}

	return result, nil
}