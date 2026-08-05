package scanner

import (
	"cfgscan/internal/config"
	"cfgscan/internal/issue"
	"cfgscan/internal/rules"
)

type Scanner struct {
	checks []rules.Rule
}

func (s Scanner) Scan(cfg config.Config) []issue.Issue {

	var issues []issue.Issue

	for _, check := range s.checks {

		issues = append(issues, check.Check(cfg)...)

	}
	return issues
}

func New() Scanner {
	return Scanner{
		checks: []rules.Rule{
			rules.DebugRule{},
			rules.UnsafeHashAlgorithmRule{},
			rules.PasswordRule{},
			rules.TLSRule{},
			rules.OpenBindAddressRule{},
		},
	}
}
