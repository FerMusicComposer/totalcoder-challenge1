package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FerMusicComposer/totalcoder-challenge1/models"
	"github.com/FerMusicComposer/totalcoder-challenge1/utils"
)

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

func jsonResponse(w http.ResponseWriter, resp FetchResponse) {
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(utils.Internal.From("Handlers/Utils/jsonResponse => error marshalling data", err).Err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
}

func successResp(records []models.Record) FetchResponse {
	return FetchResponse{
		Code:    0,
		Msg:     "Success",
		Records: records,
	}
}

func invalidDateResp() FetchResponse {
	return FetchResponse{
		Code:    utils.InvalidDate.Code,
		Msg:     utils.InvalidDate.Error(),
		Records: nil,
	}
}

func badRequestResp() FetchResponse {
	return FetchResponse{
		Code:    utils.BadRequest.Code,
		Msg:     utils.BadRequest.Error(),
		Records: nil,
	}
}

func notAllowedResp() FetchResponse {
	return FetchResponse{
		Code:    utils.NotAllowed.Code,
		Msg:     utils.NotAllowed.Error(),
		Records: nil,
	}
}

func noRecordsFoundResp() FetchResponse {
	return FetchResponse{
		Code:    utils.NotFound.Code,
		Msg:     utils.NotFound.Error(),
		Records: nil,
	}
}

func internalErrorResp() FetchResponse {
	return FetchResponse{
		Code:    utils.Internal.Code,
		Msg:     utils.Internal.Error(),
		Records: nil,
	}
}

func invalidPayloadResp() FetchResponse {
	return FetchResponse{
		Code:    utils.InvalidPayload.Code,
		Msg:     utils.InvalidPayload.Error(),
		Records: nil,
	}
}
