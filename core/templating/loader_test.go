package templating

import (
	"path"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Ein-Framework/Ein-Framework/core/domain"
	"github.com/Ein-Framework/Ein-Framework/core/services"
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"github.com/Ein-Framework/Ein-Framework/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConfig() *config.Config {
	return &config.Config{
		FrameworkRoot: ".",
		TemplatesDir:  "./_testdata",
		PluginsDir:    ".",
	}
}

func CreateManagerService(t *testing.T, config *config.Config) *TemplatingManager {
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

func TestLoadTemplateFile(t *testing.T) {
	config := CreateConfig()
	manager := CreateManagerService(t, config)

	template := manager.ReadTemplate(path.Join(config.TemplatesDir, "test.yaml"))
	t.Log(template)
}
