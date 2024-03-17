package project

import (
	"context"

	lsopenapi "github.com/hsinhoyeh/go-labelstudio/api/go-openapiv2"
	user "github.com/hsinhoyeh/go-labelstudio/pkg/user"
)

type ProjectImportData ProjectImportDataData

type ProjectImportDataData struct {
	Image string `json:"image,omitempty"`
	Text  string `json:"text,omitempty"`
}

// CreateProjectImport import url to the project for labelling.
// limitation: we now only support one task at a time.
// FIXME(hsiny): allow batch import
func (p *ProjectService) CreateProjectImport(ctx context.Context, pid int32, data ProjectImportData) (*lsopenapi.TaskCreationResponse, error) {

	importTaskCreated, _, err := p.client.OpenAPIClient().ImportApi.ApiProjectsImportCreate(
		user.NewUserService(p.client).WithAccessToken(ctx),
		lsopenapi.ImportApi{
			Data: data,
		},
		pid,
	)
	if err != nil {
		return nil, err
	}
	return &importTaskCreated, nil

}
