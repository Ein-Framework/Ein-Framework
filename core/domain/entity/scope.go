package entity

type AssetType string

const (
	Domain     AssetType = "domain"
	AndroidApp AssetType = "android"
	IOSApp     AssetType = "ios"
	ASN        AssetType = "asn"
	Wildcard   AssetType = "wildcard"
	SourceCode AssetType = "source"
)

type Asset struct {
	ID    uint      `json:"id"`
	Type  AssetType `json:"type"`
	Value string    `json:"value"`
}

type Vulnerability struct {
	Description string `json:"description"`
}

type Scope struct {
	ID                     uint            `json:"id"`
	InScope                []Asset         `json:"inScope"`
	OutScope               []Asset         `json:"outScope"`
	ClientRecommendation   string          `json:"clientRecommandations"`
	AllowedVulnerabilities []Vulnerability `json:"allowedVulnerabilities"`
	RefusedVulnerabilities []Vulnerability `json:"refusedVulnerabilities"`
}
