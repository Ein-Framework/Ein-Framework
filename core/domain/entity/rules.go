package entity

type HeaderInjectionType string

const (
	HardcodedHeader HeaderInjectionType = "hardcoded"
	PerfixedHeader  HeaderInjectionType = "prefix"
)

type HeaderInjection struct {
	HeaderName   string              `json:"name"`
	Type         HeaderInjectionType `json:"type"`
	HeaderPrefix string              `json:"prefix"`
	HeaderValue  string              `json:"value"`
}

type TestCredentials struct {
	Credentials string   `json:"credentials"`
	Domains     []string `json:"domains"`
	Description string   `json:"description"`
}

type EngagementRules struct {
	ID                 uint             `json:"id"`
	FullDescription    string           `json:"fullDescription"` // Display as markdown
	RateLimitPerSecond int              `json:"rateLimitPerSecond"`
	Threads            int              `json:"threads"`
	HeaderUse          *HeaderInjection `json:"headerInjection"`
	TestCredentials    *TestCredentials `json:"testCredentials"`
}
