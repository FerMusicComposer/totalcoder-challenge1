package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/FerMusicComposer/totalcoder-challenge1/db"
	"github.com/FerMusicComposer/totalcoder-challenge1/utils"
)

type RecordHandler struct {
	recordStore db.RecordStore
}

func NewRecordHandler(recordStore db.RecordStore) *RecordHandler {
	return &RecordHandler{
		recordStore: recordStore,
	}
}

func (rh *RecordHandler) HandleGetRecords(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println(utils.NotAllowed.From("HandleGetRecords => method not allowed", nil).Err)
		jsonResponse(w, notAllowedResp())
		return
	}

	var req FetchRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(utils.BadRequest.From("HandleGetRecords => error reading request body", err).Err)
		jsonResponse(w, badRequestResp())
		return
	}

	if err := json.Unmarshal(body, &req); err != nil {
		fmt.Println(utils.BadRequest.From("HandleGetRecords => error parsing request body", err).Err)
		jsonResponse(w, badRequestResp())
		return
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		fmt.Println(utils.InvalidDate.From("HandleGetRecords => error parsing start date", err).Err)
		jsonResponse(w, invalidDateResp())
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		fmt.Println(utils.InvalidDate.From("HandleGetRecords => error parsing end date", err).Err)
		jsonResponse(w, invalidDateResp())
		return
	}

	records, err := rh.recordStore.GetRecordsByFilter(r.Context(), startDate, endDate, req.MinCount, req.MaxCount)
	if err != nil {
		fmt.Println(utils.Internal.From("HandleGetRecords => error while obtaining records", err).Err)
		jsonResponse(w, internalErrorResp())
		return
	}

	if len(records) == 0 {
		fmt.Println(utils.NotFound.From("HandleGetRecords => no records found", nil).Err)
		jsonResponse(w, noRecordsFoundResp())
		return
	}

	fmt.Println("records received by handler:", records)
	jsonResponse(w, successResp(records))
}
