package models

type Response struct {
	GroupsResponse []GroupResponse `json:"groups"`
}

type Module struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type GroupResponse struct {
	Group   string   `json:"group"`
	Modules []Module `json:"modules"`
}

type SingleModuleResponse struct {
	Modules []Module `json:"modules"`
}
