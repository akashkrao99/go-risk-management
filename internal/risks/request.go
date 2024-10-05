package risks

import (
	"errors"
	"fmt"
)

const (
	MIN_TITLE_LEN = 5
	MIN_DESC_LEN  = 10
)

type CreateRiskReq struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      RiskStatus `json:"status"`
}

func (r *CreateRiskReq) isValid() error {
	if len(r.Title) < MIN_TITLE_LEN {
		return errors.New(fmt.Sprintf("Title must have at least %v characters", MIN_TITLE_LEN))
	}
	if len(r.Description) < MIN_DESC_LEN {
		return errors.New(fmt.Sprintf("Description must have at least %v characters", MIN_DESC_LEN))
	}
	if !isValidStatus(r.Status) {
		return errors.New("Invalid Status: " + string(r.Status))
	}
	return nil
}
