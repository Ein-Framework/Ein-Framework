package entity

import "gorm.io/gorm"

type HeaderInjectionType string

const (
	HardcodedHeader HeaderInjectionType = "hardcoded"
	PerfixedHeader  HeaderInjectionType = "prefix"
)

type HeaderInjection struct {
	gorm.Model
	HeaderName   string              `json:"name"`
	Type         HeaderInjectionType `json:"type" gorm:"type:text"`
	HeaderPrefix string              `json:"prefix"`
	HeaderValue  string              `json:"value"`
}

type TestCredentials struct {
	gorm.Model
	Credentials string   `json:"credentials"`
	Domains     []string `json:"domains" gorm:"type:text"`
	Description string   `json:"description"`
}

type EngagementRules struct {
	gorm.Model
	FullDescription    string          `json:"fullDescription"` // Display as markdown
	RateLimitPerSecond int             `json:"rateLimitPerSecond"`
	Threads            int             `json:"threads"`
	HeaderUse          HeaderInjection `json:"headerInjection" gorm:"foreignkey:HeaderUseID;association_foreignkey:ID;"`
	HeaderUseID        *uint           `json:"-"`
	TestCredentials    TestCredentials `json:"testCredentials" gorm:"foreignkey:TestCredentialsID;association_foreignkey:ID;"`
	TestCredentialsID  *uint           `json:"-"`
}
