package resolver

import (
	"context"
	"sadewa-portfolio-svc/model"
	"sadewa-portfolio-svc/repository"
)

type queryResolver struct{}
type mutationResolver struct{}

var mstPortfolioRepo = repository.MstPortfolioRepository{}

func (r *queryResolver) MstPortfolio(ctx context.Context) ([]*model.MstPortfolio, error) {
	return mstPortfolioRepo.GetAllMstPortfolios()
}

func (r *mutationResolver) CreateMstPortfolio(ctx context.Context, title string, description string) (*model.MstPortfolio, error) {
	return mstPortfolioRepo.CreateMstPortfolio(title, description)
}
