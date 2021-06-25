package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func main() {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewFileTokenStore("data.db"))

	clientStore := store.NewClientStore()
	clientStore.Set("093452", &models.Client{
		ID:     "093452",
		Secret: "824102",
		Domain: "http://0.0.0.0",
	})
	manager.MapClientStorage(clientStore)

	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)

	router := gin.Default()
	router.GET("/token", ginserver.HandleTokenRequest)

	api := router.Group("/api")
	{
		api.Use(ginserver.HandleTokenVerify())
		api.GET("/authorize", func(c *gin.Context) {
			ti, exists := c.Get(ginserver.DefaultConfig.TokenKey)
			if exists {
				c.JSON(http.StatusOK, ti)
				return
			}
			c.String(http.StatusBadRequest, "not found")
		})
	}

	router.Run(":9096")
}
