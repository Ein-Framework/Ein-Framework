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
	ID    uint
	Type  AssetType
	Value string
}

type Vulnerability struct {
	Description string
}

type Scope struct {
	ID                   uint
	InScope              []Asset
	OutScope             []Asset
	ClientRecommendation string
	Vulnerabilities      []Vulnerability
}
