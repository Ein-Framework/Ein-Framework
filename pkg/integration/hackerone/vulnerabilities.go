package hackerone

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Attributes struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	ExternalID  string `json:"external_id"`
}

type Weakness struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

type VulnerabilitiesData struct {
	Data []Weakness `json:"data"`
}

func (integration *Integration) Vulnerabilities(handle string) ([]Weakness, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.hackerone.com/v1/hackers/programs/%s/weaknesses", handle), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(integration.h1Username, integration.h1ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("program not found")
	} else if resp.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("invalid credentials for Hackerone 2")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var data VulnerabilitiesData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data.Data, nil
}
