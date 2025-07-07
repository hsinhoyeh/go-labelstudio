package invite

import (
	"context"
	"os"

	log "github.com/golang/glog"
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
	debugEnabled := os.Getenv("LABEL_STUDIO_DEBUG") != ""
	
	if debugEnabled {
		log.Info("[DEBUG] GetInvitationToken called")
	}
	
	userCtx := user.NewUserService(i.client).WithAccessToken(ctx)
	if debugEnabled {
		log.Info("[DEBUG] Created user context for API call")
	}
	
	orgInvitation, response, err := i.client.OpenAPIClient().InvitesApi.ApiInviteRead(userCtx)
	if err != nil {
		if debugEnabled {
			log.Error("[DEBUG] ApiInviteRead failed: %+v", err)
			if response != nil {
				log.Error("[DEBUG] Response status: %s", response.Status)
			}
		}
		return "", err
	}
	
	if debugEnabled {
		log.Info("[DEBUG] Successfully retrieved invitation token: %s", orgInvitation.Token)
		if response != nil {
			log.Info("[DEBUG] Response status: %s", response.Status)
		}
	}
	
	return orgInvitation.Token, nil
}
