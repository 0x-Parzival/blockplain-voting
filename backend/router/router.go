package handler

import (
    "encoding/json"
    "net/http"
    "sync"
)

var (
    blockplain = make(map[string][2]int) // region → [A count, B count]
    voteLock   sync.Mutex
)

type VoteRequest struct {
    Region string `json:"region"`
    Choice string `json:"choice"` // "A" or "B"
}

func VoteHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    var req VoteRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    if req.Choice != "A" && req.Choice != "B" {
        http.Error(w, "Invalid vote choice", http.StatusBadRequest)
        return
    }

    voteLock.Lock()
    votes := blockplain[req.Region]
    if req.Choice == "A" {
        votes[0]++
    } else {
        votes[1]++
    }
    blockplain[req.Region] = votes
    voteLock.Unlock()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Vote recorded for " + req.Region,
    })
}
