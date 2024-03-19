package invite

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	lstestutil "github.com/hsinhoyeh/go-labelstudio/pkg/http/testutil"
	"github.com/hsinhoyeh/go-labelstudio/pkg/login"
)

func TestGetInvitationToken(t *testing.T) {
	c, err := lstestutil.NewTestClient()
	assert.NoError(t, err)

	ctx := context.Background()
	assert.NoError(t, login.NewLoginService(c).DefaultLogin(ctx))

	inviteService := NewInviteService(c)
	token, err := inviteService.GetInvitationToken(ctx)
	assert.NoError(t, err)

	assert.True(t, len(token) > 0)
	fmt.Printf("token:%s\n", token)
}
