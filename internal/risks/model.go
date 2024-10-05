package risks

import (
	"errors"
	"fmt"
	"strings"
)

type RiskStatus string

const (
	STATUS_OPEN          RiskStatus = "open"
	STATUS_CLOSED        RiskStatus = "closed"
	STATUS_INVESTIGATING RiskStatus = "investigating"
	STATUS_ACCEPTED      RiskStatus = "accepted"
)

var allowedRiskStatuses []RiskStatus = []RiskStatus{STATUS_OPEN, STATUS_CLOSED, STATUS_INVESTIGATING, STATUS_ACCEPTED}

type Risk struct {
	Id          string     `json:"id" dynamodbav:"id" gorm:"column:id;primaryKey"`
	Title       string     `json:"title" dynamodbav:"title" gorm:"column:title"`
	Description string     `json:"description" dynamodbav:"description" gorm:"column:description"`
	Status      RiskStatus `json:"status" dynamodbav:"status" gorm:"column:status"`
	CreatedAt   int64      `json:"created_at" dynamodbav:"created_at" gorm:"column:created_at"`
}

func (r *Risk) sanitize() {
	r.Title = strings.TrimSpace(r.Title)
	r.Description = strings.TrimSpace(r.Description)
}

func (r *Risk) isValid() error {
	if len(r.Title) < MIN_TITLE_LEN {
		return errors.New(fmt.Sprintf("Title must have at least %v characters", MIN_TITLE_LEN))
	}
	if len(r.Description) < MIN_DESC_LEN {
		return errors.New(fmt.Sprintf("Description must have at least %v characters", MIN_DESC_LEN))
	}
	if !isValidStatus(r.Status) {
		return errors.New("Invalid Status: " + string(r.Status))
	}
	if len(r.Id) < 1 {
		return errors.New("Invalid Id: " + r.Id)
	}
	if r.CreatedAt == 0 {
		return errors.New("Creation Timestamp missing")
	}

	return nil
}
