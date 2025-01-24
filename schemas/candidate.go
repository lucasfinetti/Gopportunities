package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	Name      string    `gorm:"type:varchar(255);not null"`
	BirthDate time.Time `gorm:"type:varchar(255);not null"`
	Openings  []Opening `gorm:"many2many:candidate_opening;"`
}

type CandidateResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	Name      string    `json:"name"`
	BirthDate string    `json:"birth_date"`
}
