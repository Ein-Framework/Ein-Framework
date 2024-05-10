package entity

type AssessmentStage struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Completed    bool     `json:"completed"`
	Keywords     []string `json:"keywords"`
	Link         string   `json:"link"`
	Tasks        []Task   `json:"tasks"`
	AssessmentID uint     `json:"assessmentId"`
}

type AssessmentType string

const (
	VDP AssessmentType = "vdp"
	BB  AssessmentType = "bb"
)

type Attachement struct {
	Type string `json:"type"`
	Link string `json:"link"`
}

type Report struct {
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Attachements []Attachement `json:"attachements"`
	Severity     uint          `json:"severity"`
}

type Assessment struct {
	ID              uint
	Name            string          `json:"name"`
	Type            AssessmentType  `json:"type"`
	Scope           Scope           `json:"scope"`
	Assets          []Asset         `json:"assets"`
	Stage           AssessmentStage `json:"assessmentStage"`
	EngagementRules EngagementRules `json:"engagementRules"`
	Jobs            []Job           `json:"jobs"`
	Reports         []string        `json:"reports"`
}

func New() {
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
