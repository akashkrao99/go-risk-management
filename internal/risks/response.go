package risks

import (
	myhttp "github.com/akashkrao99/go-sample-http/pkg/http"
)

type CreateRiskRes struct {
	ResponseJson myhttp.ResponseJson `json:"response"`
	Risk         *Risk `json:"risk"`
}

type GetRiskByIdRes struct {
	ResponseJson myhttp.ResponseJson `json:"response"`
	Risk         *Risk `json:"risk"`
}

type GetRisksRes struct {
	ResponseJson myhttp.ResponseJson `json:"response"`
	Risks        []*Risk `json:"risks"`
}
