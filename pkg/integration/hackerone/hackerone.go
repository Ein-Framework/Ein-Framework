package hackerone

type Integration struct {
	h1Username string
	h1ApiKey   string
}

func New(h1Username string, h1ApiKey string) *Integration {
	return &Integration{
		h1Username,
		h1ApiKey,
	}
}
