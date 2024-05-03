package assessment

type HeaderInjectionType string

const (
	HardcodedHeader HeaderInjectionType = "hardcoded"
	PerfixedHeader  HeaderInjectionType = "prefix"
)

type HeaderInjection struct {
	HeaderName   string
	Type         HeaderInjectionType
	HeaderPrefix string
	HeaderValue  string
}

type TestCredentials struct {
	Credentials string
	Domain      []string
	Description string
}

type EngagementRules struct {
	FullDescription    string // Display as markdown
	RateLimitPerSecond int
	Threads            int
	HeaderUse          *HeaderInjection
	TestCredentials    *TestCredentials
}