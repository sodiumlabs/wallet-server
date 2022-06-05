package apis

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	oauth1tw "github.com/dghubble/oauth1/twitter"
	"github.com/gin-gonic/gin"
	"github.com/sodiumlabs/wallet-server/internal/pkg/dao/model"
	"github.com/sodiumlabs/wallet-server/internal/pkg/db"
	"gorm.io/gorm"
)

type OauthTokenResponse struct {
	AuthURL string `json:"auth_url" validate:"required"`
}

type OauthRequest struct {
	OauthCallback string `json:"oauth_callback" query:"oauth_callback" validate:"required"`
}

func RequestTwtterLoginURL(c *gin.Context, params *OauthRequest) (*OauthTokenResponse, error) {
	config := oauth1.Config{
		ConsumerKey:    os.Getenv("twconsumerKey"),
		ConsumerSecret: os.Getenv("twconsumerSecret"),
		CallbackURL:    params.OauthCallback,
		Endpoint:       oauth1tw.AuthorizeEndpoint,
	}

	requestToken, _, err := config.RequestToken()

	if err != nil {
		return nil, err
	}

	authorizationURL, err := config.AuthorizationURL(requestToken)

	if err != nil {
		return nil, err
	}

	return &OauthTokenResponse{
		AuthURL: authorizationURL.String(),
	}, nil
}

type OauthVerifierRequest struct {
	OauthToken string `json:"oauth_token" validate:"required"`
	Verifier   string `json:"verifier" validate:"required"`
}

type OauthVerifierResponse struct {
	JwtToken string `json:"jwt_token" validate:"required"`
}

// TypeName implements openapi.Typer interface for Fruit.
func (f *OauthVerifierRequest) TypeName() string { return "OauthVerifierRequestInput" }

func Auth(c *gin.Context, q *OauthVerifierRequest) (*OauthVerifierResponse, error) {
	config := oauth1.Config{
		ConsumerKey:    os.Getenv("twconsumerKey"),
		ConsumerSecret: os.Getenv("twconsumerSecret"),
		CallbackURL:    "",
		Endpoint:       oauth1tw.AuthorizeEndpoint,
	}
	// Twitter ignores the oauth_signature on the access token request. The user
	// to which the request (temporary) token corresponds is already known on the
	// server. The request for a request token earlier was validated signed by
	// the consumer. Consumer applications can avoid keeping request token state
	// between authorization granting and callback handling.
	accessToken, accessSecret, err := config.AccessToken(q.OauthToken, "secret does not matter", q.Verifier)
	if err != nil {
		return nil, err
	}

	token := oauth1.NewToken(accessToken, accessSecret)

	// httpClient will automatically authorize http.Request's
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	twuser, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{
		IncludeEmail: twitter.Bool(true),
	})

	if err != nil {
		return nil, err
	}

	// check login
	utw := &model.UserTw{}
	u := &model.User{}
	registe := false
	if err := db.DB().First(utw, "tw_id = ?", twuser.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			registe = true
			if u, utw, err = model.RegisteUserByTw(db.DB(), twuser); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	if u.Id == 0 {
		if err := db.DB().First(u, utw.UserId).Error; err != nil {
			return nil, err
		}
	}

	if registe {
		if err := InitWalletOwner(u.Id, db.DB()); err != nil {
			return nil, err
		}
	}

	jwt := NewJWT()
	jwtToken, err := jwt.CreateJWTTokenByUser(u)

	if err != nil {
		return nil, err
	}

	return &OauthVerifierResponse{
		JwtToken: jwtToken,
	}, nil
}
