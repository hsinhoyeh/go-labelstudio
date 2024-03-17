package project

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/antihax/optional"
	lsopenapi "github.com/hsinhoyeh/go-labelstudio/api/go-openapiv2"
	user "github.com/hsinhoyeh/go-labelstudio/pkg/user"
)

func (p *ProjectService) CreateProjectExport(ctx context.Context, pid int32) (*lsopenapi.ExportCreate, error) {
	exportCreated, _, err := p.client.OpenAPIClient().ExportApi.ApiProjectsExportsCreate(
		user.NewUserService(p.client).WithAccessToken(ctx),
		lsopenapi.ExportCreate{
			Title: fmt.Sprintf("auto export creation at %s", time.Now()),
			//ConvertedFormats: []lsopenapi.ConvertedFormat{
			//	lsopenapi.ConvertedFormat{
			//		ExportType: "JSON",
			//	},
			//},
		},
		pid,
	)
	if err != nil {
		return nil, err
	}
	return &exportCreated, nil
}

func (p *ProjectService) GetProjectExport(ctx context.Context, pid int32, exportId int32) (*lsopenapi.Export, error) {
	exported, _, err := p.client.OpenAPIClient().ExportApi.ApiProjectsExportsRead(
		user.NewUserService(p.client).WithAccessToken(ctx),
		pid,
		fmt.Sprintf("%d", exportId),
	)
	if err != nil {
		return nil, err
	}
	return &exported, nil
}

func (p *ProjectService) DownloadExport(ctx context.Context, pid int32, exportId int32, dst io.Writer) error {
	resp, err := p.client.OpenAPIClient().ExportApi.ApiProjectsExportsDownloadRead(
		user.NewUserService(p.client).WithAccessToken(ctx),
		pid,
		fmt.Sprintf("%d", exportId),
		&lsopenapi.ExportApiApiProjectsExportsDownloadReadOpts{
			ExportType: optional.NewString("JSON"),
		},
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	written, err := io.Copy(dst, resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%d bytes downloaded\n", written)
	return nil
}
