package model

type Mutation struct {

}

type NewMstPortfolio struct {
	ID	Int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	BackendStack string `json:"backend_stack"`
	FrontendStack string `json:"frontend_stack"`
	DatabaseStack string `json:"database_stack"`
	DeploymentStack string `json:"deployment_stack"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"creted_by"`
	UpdatedAt string `json:"updated_at"`
	IsActive bool `json:"is_active"`
}

type Query struct {

}

type MstPortfolio struct{
	NewMstPortfolio
}