package handler

import (
	"encoding/json"
	"net/http"
	"sync"
)

// Global variable for storing voting data
var (
	blockplain = make(map[string][2]int) // region → [A count, B count]
	voteLock   sync.Mutex
)

// VoteRequest represents a request to vote for a region
type VoteRequest struct {
	Region string `json:"region"`
	Choice string `json:"choice"` // "A" or "B"
}

// VoteHandler records a vote for a specific region
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
	defer voteLock.Unlock()

	// Update the vote count for the selected region and choice
	votes := blockplain[req.Region]
	if req.Choice == "A" {
		votes[0]++
	} else {
		votes[1]++
	}
	blockplain[req.Region] = votes

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Vote recorded for " + req.Region,
	})
}

// ResultsHandler returns the overall and per-region results
func ResultsHandler(w http.ResponseWriter, r *http.Request) {
	voteLock.Lock()
	defer voteLock.Unlock()

	// Prepare the result structure
	results := make(map[string]map[string]int)
	totalVotes := map[string]int{"A": 0, "B": 0}

	// Collect per-region results
	for region, votes := range blockplain {
		results[region] = map[string]int{
			"A": votes[0],
			"B": votes[1],
		}
		totalVotes["A"] += votes[0]
		totalVotes["B"] += votes[1]
	}

	// Add overall vote count
	results["overall"] = totalVotes

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
