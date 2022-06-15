package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/rs/cors"
	"github.com/sodiumlabs/wallet-server/internal/pkg/apis"
	"github.com/sodiumlabs/wallet-server/internal/pkg/apis/ethapis"
	waws "github.com/sodiumlabs/wallet-server/internal/pkg/aws"
	"github.com/sodiumlabs/wallet-server/internal/pkg/dao/model"
	"github.com/sodiumlabs/wallet-server/internal/pkg/db"
	"github.com/urfave/cli/v2"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

func Serve(c *cli.Context) error {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	err := xinit(c)

	if err != nil {
		return err
	}

	// 记录到文件。
	logf, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(logf)

	g := gin.Default()
	g.Use(gin.Recovery())
	g.Use(func(c *gin.Context) {
		cors.New(cors.Options{
			AllowOriginFunc: func(origin string) bool {
				return true
			},
			AllowedHeaders: []string{
				"Authorization",
				"Content-Type",
			},
			AllowedMethods: []string{
				"GET",
				"POST",
				"OPTIONS",
			},
			AllowCredentials: true,
			Debug:            false,
		}).HandlerFunc(c.Writer, c.Request)
	})

	f := fizz.NewFromEngine(g)

	if err := waws.Init(); err != nil {
		return err
	}

	f.Group("/auth", "auth", "auth module").
		GET("/tw_auth_url", nil, tonic.Handler(apis.RequestTwtterLoginURL, 200)).
		POST("/tw_auth", nil, tonic.Handler(apis.Auth, 200)).
		POST("/phone_auth", nil, tonic.Handler(apis.PhoneAuth, 200))

	f.Group("/wallet", "wallet", "wallet", apis.JTWMiddware).
		GET("/info", nil, tonic.Handler(apis.WalletInfo, 200)).
		GET("/pay", nil, tonic.Handler(apis.CreateCheckoutSession, 200)).
		GET("/signLog", nil, tonic.Handler(apis.SignLogApi, 200))

	f.Group("/eth", "eth", "eth", apis.JTWMiddware).
		POST("/eth_call", nil, tonic.Handler(ethapis.EthCall, 200)).
		POST("/eth_tx", nil, tonic.Handler(ethapis.EthTx, 200)).
		POST("/eth_sign", nil, tonic.Handler(ethapis.EthSign, 200))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: f,
	}

	go model.WalletDeployPool(db.DB())
	go model.StripePaySyncPool(db.DB())

	infos := &openapi.Info{
		Title:       "Wallet Server",
		Description: `Wallet Reset Server`,
		Version:     "1.0.0",
	}
	f.GET("/openapi.json", nil, f.OpenAPI(infos, "json"))

	return srv.ListenAndServe()
}
