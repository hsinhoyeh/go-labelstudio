package user

import (
	"context"
	"errors"

	"github.com/PuerkitoBio/goquery"
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
	accountUrl, err := lshttp.JoinURL(u.client.HostURL(), "/user/account")
	if err != nil {
		return "", err
	}
	accessToken, err := retrieveAccessToken(ctx, u.client, accountUrl)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (u *UserService) WithAccessToken(ctx context.Context) context.Context {
	token, err := u.GetUserToken(ctx)
	if err != nil {
		return ctx
	}
	return context.WithValue(ctx, lsopenapi.ContextAccessToken, token)
}

func retrieveAccessToken(ctx context.Context, client *lshttp.Client, url string) (string, error) {
	val, found, err := lsgoquery.ParseHTML(ctx, client, url, accessTokenParser)
	if err != nil {
		return "", err
	}
	if !found {
		return "", errors.New("found nothing with parser")
	}
	return val, nil
}

func accessTokenParser(doc *goquery.Document) (string, bool) {
	query := "[name=access_token]"
	return doc.Find(query).First().Attr("value")
}
