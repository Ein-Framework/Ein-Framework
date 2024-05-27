package hackerone_test

import (
	"testing"

	"github.com/Ein-Framework/Ein-Framework/pkg/integration/hackerone"
	"github.com/stretchr/testify/assert"
)

const username = "f1rel1ghts"
const apiKey = "cFyNuUX7tBRp42ax3N1Oz1E0RyoAwed0Qmf2KG5SGs0="

func TestProgramDetails(t *testing.T) {
	// Valid program handle
	program, err := hackerone.New(username, apiKey).ProgramDetails("zomato")
	assert.NoError(t, err)
	assert.NotNil(t, program)
	assert.Equal(t, "Zomato", program.Attributes.Name)

	// Invalid program handle
	program, err = hackerone.New(username, apiKey).ProgramDetails("invalid_handle")
	assert.Error(t, err)
	assert.Nil(t, program)
}

func TestWeaknesses(t *testing.T) {
	// Valid program handle
	weaknesses, err := hackerone.New(username, apiKey).Vulnerabilities("zomato")
	assert.NoError(t, err)
	assert.NotNil(t, weaknesses)
	assert.NotEmpty(t, weaknesses)

	// Invalid program handle
	weaknesses, err = hackerone.New(username, apiKey).Vulnerabilities("invalid_handle")
	assert.Error(t, err)
	assert.Nil(t, weaknesses)
}
