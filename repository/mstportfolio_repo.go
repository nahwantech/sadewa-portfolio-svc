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

func (repo *MstPortfolioRepository) CreateMstPortfolio(title string, description string) (*model.MstPortfolio, error) {
	mstportfolio := &model.MstPortfolio{Title: title, Description: description}
	err := model.DB.Create(mstportfolio).Error
	return mstportfolio, err
}
