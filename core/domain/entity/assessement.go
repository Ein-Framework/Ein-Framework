package entity

type AssessmentStage struct {
	Name         string
	Description  string
	Completed    string
	Keywords     []string // Alternative names for the stage
	Link         string
	Tasks        []Task
	AssessmentID uint
}

type AssessmentType string

const (
	VDP AssessmentType = "vdp"
	BB  AssessmentType = "bb"
)

type Assessment struct {
	ID              uint
	Name            string
	Type            AssessmentType
	Scope           Scope
	Assets          []Asset
	Stage           AssessmentStage
	EngagementRules EngagementRules
	Jobs            []Job
}

func New() {
	// Setup stages from the config
}

func ToggleStageCompletion() {
	// Mark stage completed / uncompleted
}

func SetCurrentStage() {
}

func CheckStagePlugins() {
}

func RunTasks() {
}

func ListStageTasks() {

}

func ViewStageTasksQueue() {
}
