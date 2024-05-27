package hackerone

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type AssetInfo struct {
	AssetType                  string `json:"asset_type"`
	AssetIdentifier            string `json:"asset_identifier"`
	EligibleForBounty          bool   `json:"eligible_for_bounty"`
	EligibleForSubmission      bool   `json:"eligible_for_submission"`
	Instruction                string `json:"instruction"`
	MaxSeverity                string `json:"max_severity"`
	CreatedAt                  string `json:"created_at"`
	UpdatedAt                  string `json:"updated_at"`
	ConfidentialityRequirement string `json:"confidentiality_requirement"`
	IntegrityRequirement       string `json:"integrity_requirement"`
	AvailabilityRequirement    string `json:"availability_requirement"`
}

type Info struct {
	Handle                          string      `json:"handle"`
	Name                            string      `json:"name"`
	Currency                        string      `json:"currency"`
	ProfilePicture                  string      `json:"profile_picture"`
	SubmissionState                 string      `json:"submission_state"`
	TriageActive                    interface{} `json:"triage_active"`
	State                           string      `json:"state"`
	StartedAcceptingAt              string      `json:"started_accepting_at"`
	NumberOfReportsForUser          int         `json:"number_of_reports_for_user"`
	NumberOfValidReportsForUser     int         `json:"number_of_valid_reports_for_user"`
	BountyEarnedForUser             float64     `json:"bounty_earned_for_user"`
	LastInvitationAcceptedAtForUser interface{} `json:"last_invitation_accepted_at_for_user"`
	Bookmarked                      bool        `json:"bookmarked"`
	AllowsBountySplitting           bool        `json:"allows_bounty_splitting"`
	OffersBounties                  bool        `json:"offers_bounties"`
	OpenScope                       interface{} `json:"open_scope"`
	FastPayments                    interface{} `json:"fast_payments"`
	GoldStandardSafeHarbor          interface{} `json:"gold_standard_safe_harbor"`
}

type Asset struct {
	ID   string    `json:"id"`
	Type string    `json:"type"`
	Info AssetInfo `json:"attributes"`
}

type Relationships struct {
	StructuredScopes struct {
		Type   string  `json:"type"`
		Assets []Asset `json:"data"`
	} `json:"structured_scopes"`
}

type Program struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Attributes    Info          `json:"attributes"`
	Relationships Relationships `json:"relationships"`
}

func (integration *Integration) ProgramDetails(handle string) (*Program, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.hackerone.com/v1/hackers/programs/%s", handle), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	fmt.Println(integration.h1Username, integration.h1ApiKey)
	req.SetBasicAuth(integration.h1Username, integration.h1ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("failed to fetch program details: %s", resp.Status)
	} else if resp.StatusCode == http.StatusUnauthorized {
		
		fmt.Println(resp.Status)
		return nil, errors.New("invalid credentials for Hackerone")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var program Program
	err = json.Unmarshal(body, &program)
	if err != nil {
		return nil, err
	}

	return &program, nil
}
