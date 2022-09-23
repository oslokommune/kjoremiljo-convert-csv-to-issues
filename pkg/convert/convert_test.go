package convert_test

import (
	"github.com/yngvark.com/kjoremiljo-convert-csv-to-issues/pkg/convert"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		expect string
	}{
		{
			name:   "Should work",
			expect: "Hello",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.expect, convert.Hello())
		})
	}
}
