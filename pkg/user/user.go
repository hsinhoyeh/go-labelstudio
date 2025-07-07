package user

import (
	"context"
	"errors"
	"os"

	"github.com/PuerkitoBio/goquery"
	log "github.com/golang/glog"
	lsopenapi "github.com/hsinhoyeh/go-labelstudio/api/go-openapiv2"
	lsgoquery "github.com/hsinhoyeh/go-labelstudio/pkg/goquery"
	lshttp "github.com/hsinhoyeh/go-labelstudio/pkg/http"
)

func NewUserService(client *lshttp.Client) *UserService {
	return &UserService{
		client: client,
	}
}

type UserService struct {
	client *lshttp.Client
}

func (u *UserService) GetUserToken(ctx context.Context) (string, error) {
	debugEnabled := os.Getenv("LABEL_STUDIO_DEBUG") != ""
	
	if debugEnabled {
		log.Info("[DEBUG] GetUserToken called")
	}
	
	accountUrl, err := lshttp.JoinURL(u.client.HostURL(), "/user/account")
	if err != nil {
		if debugEnabled {
			log.Error("[DEBUG] Failed to join URL for account: %+v", err)
		}
		return "", err
	}
	
	if debugEnabled {
		log.Info("[DEBUG] Account URL: %s", accountUrl)
	}
	
	accessToken, err := retrieveAccessToken(ctx, u.client, accountUrl)
	if err != nil {
		if debugEnabled {
			log.Error("[DEBUG] Failed to retrieve access token: %+v", err)
		}
		return "", err
	}
	
	if debugEnabled {
		log.Info("[DEBUG] Successfully retrieved access token: %s", accessToken)
	}
	
	return accessToken, nil
}

func (u *UserService) WithAccessToken(ctx context.Context) context.Context {
	debugEnabled := os.Getenv("LABEL_STUDIO_DEBUG") != ""
	
	if debugEnabled {
		log.Info("[DEBUG] WithAccessToken called")
	}
	
	token, err := u.GetUserToken(ctx)
	if err != nil {
		if debugEnabled {
			log.Error("[DEBUG] Failed to get user token for context: %+v", err)
		}
		return ctx
	}
	
	if debugEnabled {
		log.Info("[DEBUG] Adding access token to context: %s", token)
	}
	
	return context.WithValue(ctx, lsopenapi.ContextAccessToken, token)
}

func retrieveAccessToken(ctx context.Context, client *lshttp.Client, url string) (string, error) {
	debugEnabled := os.Getenv("LABEL_STUDIO_DEBUG") != ""
	
	if debugEnabled {
		log.Info("[DEBUG] Retrieving access token from URL: %s", url)
	}
	
	val, found, err := lsgoquery.ParseHTML(ctx, client, url, accessTokenParser)
	if err != nil {
		if debugEnabled {
			log.Error("[DEBUG] Failed to parse HTML for access token: %+v", err)
		}
		return "", err
	}
	if !found {
		if debugEnabled {
			log.Error("[DEBUG] Access token not found in HTML response")
		}
		return "", errors.New("found nothing with parser")
	}
	
	if debugEnabled {
		log.Info("[DEBUG] Access token found: %s", val)
	}
	
	return val, nil
}

func accessTokenParser(doc *goquery.Document) (string, bool) {
	query := "[name=access_token]"
	return doc.Find(query).First().Attr("value")
}
