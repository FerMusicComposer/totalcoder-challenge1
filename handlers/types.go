package handlers

import "github.com/FerMusicComposer/totalcoder-challenge1/models"

type FetchRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

type FetchResponse struct {
	Code    int             `json:"code"`
	Msg     string          `json:"msg"`
	Records []models.Record `json:"records"`
}
