package handler

import (
	"github.com/connormakh/pwHashApi/app/model/response"
	"github.com/connormakh/pwHashApi/app/utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatsHandlerContext_GetStats(t *testing.T) {
	setup()
	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	w := httptest.NewRecorder()
	h := StatsHandlerContext{DbSession: db}
	h.GetStats(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected error code 200, got %v", res.StatusCode)
	}
	if string(data) != utils.ToJson(response.GetStatsResponse{
		AverageTime: db.GetAverageTimeCount(),
		Total:       db.GetTotal(),
	}) {
		t.Errorf("expected Stats report, got %v", string(data))
	}
}
