package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/bitscuit/stew-chicken-bot/internal/mocks"
)

func init() {
	Client = &mocks.MockClient{}
}

func TestIsAlert(t *testing.T) {
	json, err := ioutil.ReadFile("one-alert.json")
	if err != nil {
		t.Error("Failed to get json file")
	}

	// json := `{"World_event_list":[{"metagame_event_id": "198", "timestamp": "1582857885", "metagame_event_state_name": "started"}]}`

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	got, err := isAlert("1582856026")
	if err != nil {
		t.Error(err)
	}

	wanted := "154: Hossin Enlightenment started"
	if !strings.Contains(got, wanted) {
		t.Error("Expected: \"" + wanted + "\", but got: \"" + got + "\"")
	}
}
