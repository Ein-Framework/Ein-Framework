package integration

import "github.com/Ein-Framework/Ein-Framework/pkg/integration/hackerone"

type Platforms struct {
	Hackerone *hackerone.Integration
}

func New() Platforms {
	return Platforms{}
}

func (platforms *Platforms) SetupHackerone(h1Username string, h1ApiKey string) {
	platforms.Hackerone = hackerone.New(h1Username, h1ApiKey)
}
