package gitlab_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/italolelis/watchops/internal/app/provider/gitlab"
	"github.com/stretchr/testify/assert"
)

func TestValidateToken(t *testing.T) {
	cases := []struct {
		name       string
		token      func() *http.Request
		shouldFail bool
	}{
		{
			name: "empty header",
			token: func() *http.Request {
				r, _ := http.NewRequestWithContext(
					context.Background(),
					http.MethodPost,
					"/", nil,
				)

				return r
			},
			shouldFail: true,
		},
		{
			name: "empty token",
			token: func() *http.Request {
				r, _ := http.NewRequestWithContext(
					context.Background(),
					http.MethodPost,
					"/", nil,
				)
				r.Header.Add("X-Gitlab-Token", "")

				return r
			},
			shouldFail: true,
		},
		{
			name: "invalid token",
			token: func() *http.Request {
				r, _ := http.NewRequestWithContext(
					context.Background(),
					http.MethodPost,
					"/", nil,
				)
				r.Header.Add("X-Gitlab-Token", "wrong")

				return r
			},
			shouldFail: true,
		},
		{
			name: "correct token",
			token: func() *http.Request {
				r, _ := http.NewRequestWithContext(
					context.Background(),
					http.MethodPost,
					"/", nil,
				)
				r.Header.Add("X-Gitlab-Token", "valid")

				return r
			},
			shouldFail: false,
		},
	}

	v := gitlab.NewValidator("valid")

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			err := v.Validate(c.token())
			if c.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidator_IsSupported(t *testing.T) {
	v := gitlab.NewValidator("valid")

	assert.True(t, v.IsSupported("gitlab"))
	assert.False(t, v.IsSupported("wrong"))
}
