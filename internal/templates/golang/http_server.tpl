package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"os"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	return r
}

func Start(lc fx.Lifecycle, g *gin.Engine, cfg *viper.Viper) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var port = cfg.GetString("app.server.port")
			if os.Getenv("PORT") != "" {
				port = os.Getenv("PORT")
			}
			go g.Run(fmt.Sprintf("0.0.0.0:%v", port))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("bye :)")
			return nil
		},
	})

}
