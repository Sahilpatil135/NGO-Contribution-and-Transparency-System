package models

// CauseVotesResponse is returned for upvote/downvote counters and the current user's vote.
// vote types are:
// - "up"
// - "down"
// - null (no vote)
type CauseVotesResponse struct {
	Upvotes   int     `json:"upvotes"`
	Downvotes int     `json:"downvotes"`
	MyVote    *string `json:"my_vote,omitempty"`
}

