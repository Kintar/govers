package conventionalcommit

type CommitType string

const (
	Feature CommitType = "feat"
	Fix     CommitType = "fix"
)

type Footer struct {
	Token string
	Value string
}

type ChangeDetail struct {
	CommitType  CommitType
	Scope       string
	Description string
	CommitBody  string
	CommitHash  string
}
