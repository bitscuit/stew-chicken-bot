package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bitscuit/stew-chicken-bot/internal/mocks"
)

func init() {
	Client = &mocks.MockClient{}
}

func TestIsAlert(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"one-alert.json", "154: Hossin Enlightenment started"},
		{"no-alerts.json", "No alerts"},
	}

	for _, test := range tests {
		path := filepath.Join("testdata", "ps2", test.in)
		json, err := ioutil.ReadFile(path)
		if err != nil {
			t.Error("Failed to get json file")
		}

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

		wanted := test.out
		if !strings.Contains(got, wanted) {
			t.Error("Expected: \"" + wanted + "\", but got: \"" + got + "\"")
		}
	}

}
