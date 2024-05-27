package entity

import "gorm.io/gorm"

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
	gorm.Model
	Type  AssetType `json:"type" gorm:"type:text"`
	Value string    `json:"value"`
}

type Vulnerability struct {
	gorm.Model
	Description string `json:"description"`
	ExternalID  string `json:"external_id"`
	Name        string `json:"name"`
}

type Scope struct {
	gorm.Model
	InScope                []Asset         `json:"inScope" gorm:"many2many:scope_in_assets;"`
	OutScope               []Asset         `json:"outScope" gorm:"many2many:scope_out_assets;"`
	ClientRecommendation   string          `json:"clientRecommandations"`
	AllowedVulnerabilities []Vulnerability `json:"allowedVulnerabilities" gorm:"many2many:scope_allowed_vulnerabilities;"`
	RefusedVulnerabilities []Vulnerability `json:"refusedVulnerabilities" gorm:"many2many:scope_refused_vulnerabilities;"`
}
