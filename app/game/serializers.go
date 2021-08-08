package game

import (
	"github.com/gin-gonic/gin"
	"time"
)

// Response -
// The game structure to be sent in the response when required
type Response struct {
	ExternalID string `json:"external_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	ClickScore Score  `json:"click_score"`
	Status     Status `json:"status"`
}

// Serializer -
// Holds the needed objects to perform the serialization
type Serializer struct {
	C *gin.Context
	Model
}

// Response -
// The function that builds the response
func (s *Serializer) Response() Response {
	return Response{
		ExternalID: s.ExternalID,
		CreatedAt:  s.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt:  s.UpdatedAt.UTC().Format(time.RFC3339),
		ClickScore: s.ClickScore,
		Status:     s.Status,
	}
}
