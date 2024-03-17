package project

import (
	"bytes"
	"context"
	"fmt"
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

	projService := NewProjectService(c)

	resp, err := projService.CreateProject(
		ctx,
		"test title from gotest",
		"empty desc",
		LabelConfigDefaultImageClassification,
	)
	assert.NoError(t, err)

	resp2, err := projService.GetProject(ctx, resp.Id)
	assert.NoError(t, err)
	assert.EqualValues(t, resp2, resp)

	// import data to the project created

	respn, err := projService.CreateProjectImport(ctx, resp.Id,
		ProjectImportData(
			ProjectImportDataData{
				Image: "https://example.com/somepic.jpg",
			},
		),
	)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, respn.TaskCount)

	// export data from the project
	exportCreated, err := projService.CreateProjectExport(ctx, resp.Id)
	assert.NoError(t, err)
	//assert.EqualValues(t, 1, respn.TaskCount)
	fmt.Printf("exportcretaed:%+v\n", exportCreated)

	exported, err := projService.GetProjectExport(ctx, resp.Id, exportCreated.Id)
	assert.NoError(t, err)

	//assert.EqualValues(t, exported, exportCreated)
	fmt.Printf("exported:%+v\n", exported)

	downloadByteBuffer := &bytes.Buffer{}
	assert.NoError(t, projService.DownloadExport(ctx, resp.Id, exportCreated.Id, downloadByteBuffer))
	fmt.Printf("d:%s\n", downloadByteBuffer.String())
}
