package invite

import (
	"context"

	lshttp "github.com/hsinhoyeh/go-labelstudio/pkg/http"
	"github.com/hsinhoyeh/go-labelstudio/pkg/user"
)

func NewInviteService(client *lshttp.Client) *InviteService {
	return &InviteService{
		client: client,
	}
}

type InviteService struct {
	client *lshttp.Client
}

func (i *InviteService) GetInvitationToken(ctx context.Context) (string, error) {
	orgInvitation, _, err := i.client.OpenAPIClient().InvitesApi.ApiInviteRead(
		user.NewUserService(i.client).WithAccessToken(ctx),
	)
	if err != nil {
		return "", err
	}
	return orgInvitation.Token, nil
}
