package candidate

import (
	"fmt"
	"time"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateCandidate

type CreateCandidateRequest struct {
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
}

func (r *CreateCandidateRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	return nil
}
