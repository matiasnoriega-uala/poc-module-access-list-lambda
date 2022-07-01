package models

type Request struct {
	UserEmail string `json:"userEmail"`
}

type GroupModulesRequest struct {
	Group string `json:"group"`
}
