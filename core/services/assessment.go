package services

import (
	"fmt"

	"github.com/Ein-Framework/Ein-Framework/core/domain/entity"
	"github.com/Ein-Framework/Ein-Framework/pkg/integration"
	"github.com/Ein-Framework/Ein-Framework/pkg/integration/hackerone"
	"github.com/Ein-Framework/Ein-Framework/pkg/repository"
	"gorm.io/gorm"
)

type AssessmentService struct {
	Service
	platformsIntegration integration.Platforms
}

func NewAssessmentService(ctx Context) *AssessmentService {
	repo := repository.NewGormRepository(ctx.OrmConnection.Db, ctx.Logger.Sugar())

	platformsIntegration := integration.New()
	platformsIntegration.SetupHackerone(ctx.Config.Integration.HackeroneCreds.H1Username, ctx.Config.Integration.HackeroneCreds.H1ApiKey)

	return &AssessmentService{
		Service: Service{
			repo:          repo,
			ormConnection: ctx.OrmConnection,
			logger:        ctx.Logger,
		},
		platformsIntegration: platformsIntegration,
	}
}

func (s *AssessmentService) GetAllAssessments() (*[]entity.Assessment, error) {
	var assessments []entity.Assessment

	err := s.repo.GetAll(&assessments, "Assets", "Stage", "Scope.InScope", "Scope.OutScope")
	if err != nil {
		return nil, err
	}

	fmt.Println(assessments)

	return &assessments, nil
}

func (s *AssessmentService) DeleteAssesment(id uint) error {
	return s.repo.Delete(&entity.Assessment{Model: gorm.Model{ID: id}})
}

func (s *AssessmentService) AddNewAssessment(name string, assessmentType entity.AssessmentType, scope entity.Scope, engagementRules entity.EngagementRules) (*entity.Assessment, error) {
	assessment, err := entity.NewAssessment(name, assessmentType, scope, engagementRules, s.repo)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(assessment)
	if err != nil {
		return nil, err
	}

	return assessment, nil
}

func (s *AssessmentService) DeleteAssessment(id uint) error {
	assessment := &entity.Assessment{Model: gorm.Model{ID: id}}
	return s.repo.Delete(assessment)
}

func (s *AssessmentService) UpdateAssessment(id uint, updatedAssessment *entity.Assessment) error {
	assessment, err := s.GetAssessmentById(id)
	if err != nil {
		return err
	}

	assessment.Name = updatedAssessment.Name
	assessment.Type = updatedAssessment.Type
	assessment.Scope = updatedAssessment.Scope
	assessment.Assets = updatedAssessment.Assets
	assessment.Stage = updatedAssessment.Stage
	assessment.EngagementRules = updatedAssessment.EngagementRules
	// assessment.Jobs = updatedAssessment.Jobs
	assessment.Reports = updatedAssessment.Reports

	return s.repo.Save(assessment)
}

func (s *AssessmentService) GetAssessmentById(id uint) (*entity.Assessment, error) {
	var assessment entity.Assessment
	err := s.repo.GetOneByID(&assessment, id, "Scope.InScope", "Scope.OutScope")
	if err != nil {
		return nil, err
	}
	return &assessment, nil
}

func (s *AssessmentService) AddNewAssessmentFromHackerone(programName string) (*entity.Assessment, error) {

	program, err := s.platformsIntegration.Hackerone.ProgramDetails(programName)

	if err != nil {
		return nil, err
	}

	vulnerabilities, err := s.platformsIntegration.Hackerone.Vulnerabilities(programName)

	if err != nil {
		return nil, err
	}

	assesment := s.HackeroneToAssessmentAdapter(*program, vulnerabilities)

	err = s.repo.Create(&assesment)

	if err != nil {
		return nil, err
	}

	return &assesment, nil
}

func (s *AssessmentService) HackeroneToAssessmentAdapter(program hackerone.Program, weaknesses []hackerone.Weakness) entity.Assessment {

	var assessmentType string
	if program.Attributes.OffersBounties {
		assessmentType = "Bug Bounty"
	} else {
		assessmentType = "VDP"
	}

	assets := []entity.Asset{}
	inScope := []entity.Asset{}
	outScope := []entity.Asset{}

	for _, asset := range program.Relationships.StructuredScopes.Assets {

		frameworkAsset := entity.Asset{
			Value: asset.Info.AssetIdentifier,
			Type:  hackeroneAssetTypeAdapater(asset.Info.AssetType),
		}
		assets = append(assets, frameworkAsset)
		if asset.Info.EligibleForSubmission {
			inScope = append(inScope, frameworkAsset)

		} else {
			outScope = append(outScope, frameworkAsset)
		}
	}

	vulnerabilities := []entity.Vulnerability{}

	for _, weakness := range weaknesses {
		vulnerabilities = append(vulnerabilities, entity.Vulnerability{
			Name:        weakness.Attributes.Name,
			Description: weakness.Attributes.Description,
			ExternalID:  weakness.Attributes.ExternalID,
		})
	}

	assessment := entity.Assessment{
		Name:   program.Attributes.Name,
		Type:   entity.AssessmentType(assessmentType),
		Assets: assets,
		Scope: entity.Scope{
			InScope:                inScope,
			OutScope:               outScope,
			AllowedVulnerabilities: vulnerabilities,
		},
		Stage: entity.AssessmentStage{
			Name:        entity.ReconnaissanceStage,
			Description: "",
			Completed:   false,
		},
		EngagementRules: entity.EngagementRules{
			Threads:            10,
			RateLimitPerSecond: 10,
			FullDescription:    "",
		},
		// Jobs:    []entity.Job{},
		Reports: []entity.Report{},
	}

	return assessment

}

func hackeroneAssetTypeAdapater(assetType string) entity.AssetType {

	assetTypesMap := map[string]entity.AssetType{
		"URL":                entity.Domain,
		"GOOGLE_PLAY_APP_ID": entity.AndroidApp,
		"WILDCARD":           entity.Wildcard,
		"APPLE_STORE_APP_ID": entity.IOSApp,
	}

	return assetTypesMap[assetType]
}
