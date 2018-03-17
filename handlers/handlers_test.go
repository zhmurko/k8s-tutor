package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"tutor/version"
)

func TestRouter(t *testing.T) {
	r := Router(version.BuildTime, version.Commit, version.Release)
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/status")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /status is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	res, err = http.Post(ts.URL+"/status", "text/plain", nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status code for /status is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusMethodNotAllowed)
	}

	res, err = http.Get(ts.URL + "/not-exists")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Status code for /status is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusNotFound)
	}
}
