package templating

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Ein-Framework/Ein-Framework/core/domain"
	"github.com/Ein-Framework/Ein-Framework/core/services"
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"github.com/Ein-Framework/Ein-Framework/pkg/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	TEMPLATE_TEST_FILE_NAME = "test.yaml"
)

func CreateConfig() *config.Config {
	return &config.Config{
		FrameworkRoot: ".",
		TemplatesDir:  "./_testdata",
		PluginsDir:    ".",
	}
}

func CreateManagerService(t *testing.T, config *config.Config) ITemplateManager {
	logger := log.GetLogger()

	mockdb, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: mockdb,
	}), &gorm.Config{})
	if err != nil {
		t.Fatal(err.Error())
	}
	defer mockdb.Close()

	coreServices := services.InitServices(
		&domain.ORMConnection{
			Db: gormDB,
		},
		logger,
		config,
	)
	return New(config, coreServices, logger)
}

func TestReadTemplateFile(t *testing.T) {
	config := CreateConfig()
	manager := CreateManagerService(t, config)

	template, err := manager.ReadTemplate("test.yaml")

	if err != nil {
		return
	}

	assert.Equal(t, err, nil, "There should be no errors")
	assert.Equal(t, len(template.Steps), 1, "There should be 1 step only")
	assert.Equal(t, template.Steps[0].Protocol, "http", "Expected protocol is http")
}

func TestListAllTemplates(t *testing.T) {
	config := CreateConfig()
	manager := CreateManagerService(t, config)

	templates, err := manager.ListAllAvailableTemplates()

	assert.Equal(t, err, nil, "There should be no errors")
	if err != nil {
		return
	}

	assert.Equal(t, len(templates), 1, "There should be 1 template only")
	assert.Equal(t, templates[0], TEMPLATE_TEST_FILE_NAME)
	fmt.Print(templates)
}
