package issue

type Severity string

const (
	Low    Severity = "LOW"
	Medium Severity = "MEDIUM"
	High   Severity = "HIGH"
)

type Issue struct {
	Severity       Severity
	Message        string
	Recommendation string
}
