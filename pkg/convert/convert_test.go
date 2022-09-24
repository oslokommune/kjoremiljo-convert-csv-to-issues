package convert_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yngvark.com/kjoremiljo-convert-csv-to-issues/pkg/convert"

	"github.com/stretchr/testify/assert"
)

// nolint:lll
const csv = `
Component,Feature,Skal flyttes til Github Projects,Flyttekommentar,Flyttet til Github Projects,Usertesting,Done ✔️,V1,V2,V3,Done,Documentation only,Comments
AWS account,Guardrails in control tower,x,,,,,,,,,,
Containers,CI/CD toolchain / GitOps - oppdaterar terraform repo,x,,,x,,,,,,,"Is it possible to run in Fargate, to avoid giving too much permissions for a Github Action? Or should we use wiz.io ?"
Div,IPv6,x,,,x,,,,,,,
Div,Blah blah,,,,x,,,,,,,
`

func TestCreateGithubIssues(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		expect []string
	}{
		{
			name: "Should work",
			// nolint:lll
			expect: []string{
				"gh issue create --repo 'bob/issue-repo' --label 'AWS account' --title 'Guardrails in control tower' --body ''",
				"gh issue create --repo 'bob/issue-repo' --label 'Containers' --title 'CI/CD toolchain / GitOps - oppdaterar terraform repo' --body 'Is it possible to run in Fargate, to avoid giving too much permissions for a Github Action? Or should we use wiz.io ?'",
				"gh issue create --repo 'bob/issue-repo' --label '' --title 'IPv6' --body ''",
			},
		},
	}

	for _, tc := range testCases {
		//nolint:varnamelen
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			commands, err := convert.CreateGhCommands(csv, "bob/issue-repo")
			require.NoError(t, err)

			assert.Equal(t, len(tc.expect), len(commands))

			for i, cmd := range commands {
				assert.Equal(t, tc.expect[i], cmd)
			}
		})
	}
}
