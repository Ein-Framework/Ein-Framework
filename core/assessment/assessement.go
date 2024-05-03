package assessment

type AssessmentStage struct {
	Description string
	Name        string
	Completed   string
	Keywords    []string // Alternative names for the stage
	Link        string
	Tasks       []Task
}

type AssessmentType string

const (
	VDP AssessmentType = "vdp"
	BB  AssessmentType = "bb"
)

type Assessment struct {
	Name            string
	Type            AssessmentType
	Scope           Scope
	Assets          []Asset
	Stage           AssessmentStage
	EngagementRules EngagementRules
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
