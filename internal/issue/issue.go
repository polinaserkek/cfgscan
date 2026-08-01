package issue

type Severity string

const (
	LOW    Severity = "LOW"
	MEDIUM Severity = "MEDIUM"
	HIGH   Severity = "HIGH"
)

type Issue struct {
	Severity       Severity
	Message        string
	Recommendation string
}
