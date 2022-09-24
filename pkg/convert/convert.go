package convert

import (
	"encoding/csv"
	"fmt"
	"strings"
)

// CreateGhCommands returns gh CLI commands in a string array, for instance:
// gh issue create --repo yngvark/temp --title "Some title" --body "Some body"
// gh issue create --repo yngvark/temp --title "Some title 2" --body "Some body 2"
// .
func CreateGhCommands(csvContents string, issueRepo string) ([]string, error) {
	issues, err := parseIssuesFromCsv(csvContents)
	if err != nil {
		return []string{}, err
	}

	issues = removeDivLabel(issues)
	issues = removeNotMovingToGithub(issues)
	commands := toCommands(issues, issueRepo)

	return commands, nil
}

func parseIssuesFromCsv(csvContents string) ([]Issue, error) {
	reader := csv.NewReader(strings.NewReader(csvContents))

	records, err := reader.ReadAll()
	if err != nil {
		return []Issue{}, fmt.Errorf("reading CSV: %w", err)
	}

	issues := make([]Issue, 0)

	for i, record := range records {
		if i == 0 {
			continue
		}

		moveToGithub := false
		if record[2] == "x" {
			moveToGithub = true
		}

		issue := Issue{
			Label:        record[0],
			Title:        record[1],
			Body:         record[12],
			moveToGithub: moveToGithub,
		}

		issues = append(issues, issue)
	}

	return issues, nil
}

func removeDivLabel(issues []Issue) []Issue {
	newIssues := make([]Issue, 0)

	for _, issue := range issues {
		if issue.Label == "Div" {
			issue.Label = ""
		}

		newIssues = append(newIssues, issue)
	}

	return newIssues
}

func removeNotMovingToGithub(issues []Issue) []Issue {
	filtered := make([]Issue, 0)

	for _, issue := range issues {
		if issue.moveToGithub {
			filtered = append(filtered, issue)
		}
	}

	return filtered
}

func toCommands(issues []Issue, issueRepo string) []string {
	commands := make([]string, 0)

	for _, issue := range issues {
		cmd := fmt.Sprintf("gh issue create --repo '%s' --label '%s' --title '%s' --body '%s'",
			issueRepo, issue.Label, issue.Title, issue.Body)

		commands = append(commands, cmd)
	}

	return commands
}
