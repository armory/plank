package plank

type PlankClient interface {
	// Applications
	GetApplication(string) (*Application, error)
	GetApplications() (*[]Application, error)
	CreateApplication(*Application) error

	// Permissions
	IsAdmin(string) (bool, error)
	HasAppWriteAccess(string, string) (bool, error)
	GetUser(string) (*User, error)

	// Pipelines
	GetPipeline(string, string) (*Pipeline, error)
	GetPipelines(string) ([]Pipeline, error)
	UpsertPipeline(Pipeline) error
	DeletePipeline(Pipeline) error
	DeletePipelineByName(string, string) error
	Execute(string, string) (*PipelineRef, error)

	// Tasks
	GetTask(string) (*ExecutionStatusResponse, error)
	PollTaskStatus(string) (*ExecutionStatusResponse, error)
	CreateTask(string, string, interface{}) (*TaskRefResponse, error)
}
