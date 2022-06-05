package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/sodiumlabs/wallet-server/internal/pkg/dao/model"
	"github.com/sodiumlabs/wallet-server/internal/pkg/db"
	"gorm.io/gorm"
)

type PhoneAuthRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Verifier string `json:"verifier" validate:"required"`
}

type PhoneAuthResponse struct {
	JwtToken string `json:"jwt_token" validate:"required"`
}

// TypeName implements openapi.Typer interface for Fruit.
func (f *PhoneAuthRequest) TypeName() string { return "PhoneAuthRequestInput" }

func PhoneAuth(c *gin.Context, q *PhoneAuthRequest) (*PhoneAuthResponse, error) {
	// check login
	up := &model.UserPhone{}
	u := &model.User{}
	registe := false
	if err := db.DB().First(up, "phone = ?", q.Phone).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			registe = true
			if u, up, err = model.RegisteUserByPhone(db.DB(), q.Phone); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	if u.Id == 0 {
		if err := db.DB().First(u, up.UserId).Error; err != nil {
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

	return &PhoneAuthResponse{
		JwtToken: jwtToken,
	}, nil
}
