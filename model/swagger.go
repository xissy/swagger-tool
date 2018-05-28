package model

type Swagger struct {
	Swagger string `json:"swagger"`
	Info map[string]string `json:"info"`
	Schemes []string `json:"schemes"`
	Consumes []string `json:"consumes"`
	Produces []string `json:"produces"`
	Paths map[string]*Path `json:"paths"`
	Definitions map[string]*Definition `json:"definitions"`
}

type Path interface{}

type Definition interface{}
