package apis

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sodiumlabs/wallet-server/internal/pkg/dao/model"
	"github.com/sodiumlabs/wallet-server/internal/pkg/db"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func init() {
	stripe.Key = "sk_test_51KRXxHLqvgNlwPxdqeDcqQCbTl85ZGmsLNa8sHRU45OQUAFQgVvO6ZYAvaoNR2w974Cw9N05ReZvse7XISu0fjeN00cg8ZImNG"
}

type CreateCheckoutSessionReq struct {
	SuccessCallbackURL string `json:"success_callback_url" query:"success_callback_url" validate:"required"`
	CancelCallbackURL  string `json:"cancel_callback_url" query:"cancel_callback_url" validate:"required"`
	Amount             int64  `json:"amount" query:"amount" validate:"required"`
}

type CreateCheckoutSessionRes struct {
	PaymentURL string `json:"payment_url" validate:"required"`
}

func CreateCheckoutSession(c *gin.Context, q *CreateCheckoutSessionReq) (*CreateCheckoutSessionRes, error) {
	userId := c.GetUint("user_id")

	var user model.User
	if err := db.DB().First(&user, userId).Error; err != nil {
		return nil, err
	}

	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(fmt.Sprintf("Buy MELD to %s", user.WalletAddress)),
					},
					UnitAmount: stripe.Int64(q.Amount),
				},
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String(q.SuccessCallbackURL),
		CancelURL:  stripe.String(q.CancelCallbackURL),
	}

	s, err := session.New(params)

	if err != nil {
		return nil, err
	}

	sp := model.StripePay{
		Id:                s.ID,
		Amount:            uint64(s.AmountTotal),
		UserWalletAddress: user.WalletAddress,
		Status:            s.Status,
		LatestCheckedAt:   time.Now().UTC(),
	}

	if err := db.DB().Create(&sp).Error; err != nil {
		return nil, err
	}

	return &CreateCheckoutSessionRes{
		PaymentURL: s.URL,
	}, nil
}
