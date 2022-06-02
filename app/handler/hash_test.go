package handler

import (
	"bytes"
	"encoding/json"
	"github.com/connormakh/pwHashApi/app/model/response"
	"github.com/connormakh/pwHashApi/app/utils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sync"
	"testing"
)

var db *utils.Datastore

func setup() {
	db = utils.InitializeDatastore()
}

func TestHashHandlerContext_GetHash_NotFound(t *testing.T) {
	setup()
	req := httptest.NewRequest(http.MethodGet, "/hash/1", nil)
	w := httptest.NewRecorder()
	h := HashHandlerContext{DbSession: db, Wg: nil}
	h.GetHash(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected error code 404, got %v", res.StatusCode)
	}
	if string(data) != utils.ToJson(response.NewGetHashNotFoundError()) {
		t.Errorf("expected hash not found error, got %v", string(data))
	}
}

func TestHashHandlerContext_GetHash_MalformedId(t *testing.T) {
	setup()
	req := httptest.NewRequest(http.MethodGet, "/hash/1malformed", nil)
	w := httptest.NewRecorder()
	h := HashHandlerContext{DbSession: db, Wg: nil}
	h.GetHash(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected error code 400, got %v", res.StatusCode)
	}
	if string(data) != utils.ToJson(response.NewGetHashMalformedIdError()) {
		t.Errorf("expected hash malformed id error, got %v", string(data))
	}
}

func TestHashHandlerContext_GetHash(t *testing.T) {
	setup()
	db.Increment()
	db.Insert(1, "Test")

	req := httptest.NewRequest(http.MethodGet, "/hash/1", nil)
	w := httptest.NewRecorder()
	h := HashHandlerContext{DbSession: db, Wg: nil}
	h.GetHash(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected error code 200, got %v", res.StatusCode)
	}
	if string(data) != utils.ToJson(response.GetHashResponse{PasswordHash: "Test"}) {
		t.Errorf("expected hash response, got %v", string(data))
	}
}

func TestHashHandlerContext_PostHash(t *testing.T) {
	setup()
	mcPostBody := map[string]interface{}{
		"password": "test",
	}
	body, _ := json.Marshal(mcPostBody)
	req := httptest.NewRequest(http.MethodPost, "/hash", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	h := HashHandlerContext{DbSession: db, Wg: &sync.WaitGroup{}}
	h.PostHash(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	matched, _ := regexp.MatchString(`{\"id":\d+}`, string(data))
	if !matched {
		t.Errorf("expected hash response, got %v", string(data))
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected error code 200, got %v", res.StatusCode)
	}
	//time.Sleep(5 * time.Second)
	//db.FindOne()
}
