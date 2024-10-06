package risks_test

import (
	"github.com/akashkrao99/go-sample-http/internal/risks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRiskReqIsValid(t *testing.T) {
	tests := []struct {
		name      string
		request   *risks.CreateRiskReq
		expectErr bool
		errMsg    string
	}{
		{
			name: "Valid request",
			request: &(risks.CreateRiskReq{
				Title:       "Valid Title",
				Description: "Valid description here.",
				Status:      risks.STATUS_OPEN,
			}),
			expectErr: false,
		},
		{
			name: "Title too short",
			request: &(risks.CreateRiskReq{
				Title:       "Shrt",
				Description: "Valid description here.",
				Status:      risks.STATUS_OPEN,
			}),
			expectErr: true,
			errMsg:    "title length must be between 5 to 20 characters",
		},
		{
			name: "Title too long",
			request: &(risks.CreateRiskReq{
				Title:       "This title is way too long",
				Description: "Valid description here.",
				Status:      risks.STATUS_OPEN,
			}),
			expectErr: true,
			errMsg:    "title length must be between 5 to 20 characters",
		},
		{
			name: "Description too short",
			request: &(risks.CreateRiskReq{
				Title:       "Valid Title",
				Description: "Short",
				Status:      risks.STATUS_OPEN,
			}),
			expectErr: true,
			errMsg:    "description length must be between 10 to 30 characters",
		},
		{
			name: "Description too long",
			request: &(risks.CreateRiskReq{
				Title:       "Valid Title",
				Description: "This description is way too long and exceeds the character limit.",
				Status:      risks.STATUS_OPEN,
			}),
			expectErr: true,
			errMsg:    "description length must be between 10 to 30 characters",
		},
		{
			name: "Invalid status",
			request: &(risks.CreateRiskReq{
				Title:       "Valid Title",
				Description: "Valid description here.",
				Status:      "invalid",
			}),
			expectErr: true,
			errMsg:    "Invalid Status: invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.IsValid()
			if tt.expectErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
