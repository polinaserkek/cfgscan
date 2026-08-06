package permissions

import (
	"cfgscan/internal/issue"
	"os"
)

func CheckFilePermissions(path string) ([]issue.Issue, error) {
	info, err := os.Stat(path)

	if err != nil {
		return nil, err
	}
	perms := info.Mode().Perm()

	issues := []issue.Issue{}

	if perms&0o002 != 0 {
		issues = append(issues, issue.Issue{
			Severity:       issue.MEDIUM,
			Message:        "файл доступен на запись для остальных пользователей. ",
			Recommendation: "Ограничьте права доступа к файлу.",
		})

	}
	if perms&0o020 != 0 {
		issues = append(issues, issue.Issue{
			Severity:       issue.MEDIUM,
			Message:        "конфигурационный файл доступен на запись для группы. ",
			Recommendation: "Ограничьте права доступа к файлу конфигурации.",
		})
	}
	return issues, nil
}
