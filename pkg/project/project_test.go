package project

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	lstestutil "github.com/hsinhoyeh/go-labelstudio/pkg/http/testutil"
	"github.com/hsinhoyeh/go-labelstudio/pkg/login"
)

func TestCreateProject(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
		return
	}

	c, err := lstestutil.NewTestClient()
	assert.NoError(t, err)

	ctx := context.Background()
	assert.NoError(t, login.NewLoginService(c).DefaultLogin(ctx))

	resp, err := NewProjectService(c).CreateProject(
		ctx,
		"test title from gotest",
		"empty desc",
		LabelConfigDefaultKeyPointLabels,
	)
	assert.NoError(t, err)

	resp2, err := NewProjectService(c).GetProject(ctx, resp.ID)
	assert.NoError(t, err)
	assert.EqualValues(t, resp2, resp)
}
