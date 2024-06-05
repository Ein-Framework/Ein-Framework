package requests

import "github.com/Ein-Framework/Ein-Framework/core/domain/entity"

type CreateAssessmentFromURL struct {
	ProgramName string `json:"programName"`
	Platform    string `json:"platform"`
}

type AssessmentRequest struct {
	Id              uint                   `param:"id"`
	Name            string                 `json:"name"`
	AssessmentType  entity.AssessmentType  `json:"assessmentType"`
	Scope           entity.Scope           `json:"scope"`
	EngagementRules entity.EngagementRules `json:"engagementRules"`
}

type IdParam struct {
	Id uint `param:"id"`
}
