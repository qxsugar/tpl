package http

import (
	"api-pkg/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/qxsugar/pkg/apix"
	prometheus "github.com/zsais/go-gin-prometheus"
)

func Start() {
	app := gin.Default()
	app.Use(middleware.GlobalErrorHandler())

	cnf := cors.DefaultConfig()
	cnf.AddAllowHeaders("TOKEN", "token")
	cnf.AllowAllOrigins = true
	app.Use(cors.New(cnf))

	p := prometheus.NewPrometheus("gin")
	p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		return c.FullPath()
	}
	p.Use(app)

	app.GET("/", apix.Pong) // ping

	_ = app.Run("0.0.0.0:8000")
}
