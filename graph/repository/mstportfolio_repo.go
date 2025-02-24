package repository

import (
	"sadewa-portfolio-svc/model"
)

type MstPortfolioRepository struct{}

func (repo *MstPortfolioRepository) GetAllMstPortfolios() ([]*model.MstPortfolio, error) {
	var mstportfolios []*model.MstPortfolio
	err := model.DB.Preload("MstPortfolio").Find(&mstportfolios).Error
	return mstportfolios, err
}