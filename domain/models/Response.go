package models

type Response struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	//Warnings    []string `json:"warnings"`
}
