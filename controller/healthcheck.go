package controller

import (
    "encoding/json"
	"net/http"
    
    _ "github.com/mikasoftware/mikapost-go/base/database"
)


type HealthCheckStatus struct {
    Details string
}

func HealthCheckFunc(w http.ResponseWriter, r *http.Request) {
    data := HealthCheckStatus{
        Details: "Welcome to the ComicsCantina backend API, build v0.0.001.0",
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(data)
}
