package handlers

import (
    "tutor/version"
    "encoding/json"
    "log"
    "net/http"
)

// status is a simple HTTP handler function which writes a response.
func status(w http.ResponseWriter, _ *http.Request) {
  info := struct {
       BuildTime string `json:"buildTime"`
       Commit    string `json:"commit"`
       Release   string `json:"release"`
   }{
       version.BuildTime, version.Commit, version.Release,
   }

   body, err := json.Marshal(info)
   if err != nil {
       log.Printf("Could not encode info data: %v", err)
       http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
       return
   }
   w.Header().Set("Content-Type", "application/json")
   w.Write(body)
}
