package risks

import (
	"errors"
	"fmt"
)

const (
	MIN_TITLE_LEN = 5
	MIN_DESC_LEN  = 10
	MAX_TITLE_LEN = 20
	MAX_DESC_LEN  = 30
)

type CreateRiskReq struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      RiskStatus `json:"status"`
}

func (r *CreateRiskReq) isValid() error {

	if len(r.Title) < MIN_TITLE_LEN || len(r.Title) > MAX_TITLE_LEN {
		return fmt.Errorf("title length must be between %v to %v characters", MIN_TITLE_LEN, MAX_TITLE_LEN)
	}
	if len(r.Description) < MIN_DESC_LEN || len(r.Description) > MAX_DESC_LEN {
		return fmt.Errorf("description length must be between %v to %v characters", MIN_DESC_LEN, MAX_DESC_LEN)
	}
	if !isValidStatus(r.Status) {
		return errors.New("Invalid Status: " + string(r.Status))
	}
	return nil
}
