package server

import (
	"log"
	"net/http"

	"github.com/CrocSwap/graphcache-go/views"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type APIWebServer struct {
	Views views.IViews
}

func (s *APIWebServer) Serve(prefix string, extendedApi bool) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.GET("/", func(c *gin.Context) { c.Status(http.StatusOK) })
	r.GET(prefix+"/", func(c *gin.Context) { c.Status(http.StatusOK) })
	r.GET(prefix+"/created_contracts", s.queryCreatedContracts)

	log.Println("API Serving at", prefix)
	r.Run()
}

func (s *APIWebServer) queryCreatedContracts(c *gin.Context) {
	chainId := parseChainParam(c, "chainId")
	if len(c.Errors) > 0 {
		return
	}
	resp := s.Views.QueryCreatedContracts(chainId)
	wrapDataErrResp(c, resp, nil)
}
