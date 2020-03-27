package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "net/http"
)

var (
	router = echo.New()
	jwtKey = getSecretKey()
	isProd = flag.Bool("p", false, "Run in production")
)

func main() {

	flag.Parse()

	router.Use(middleware.Static("./dist"))
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}",` +
			`"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
	}))
	router.Use(authMiddleware)
//	 router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
//	 	AllowOrigins:     []string{"http://101.53.146.105:8080"},
//	 	AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	 //	AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
//	 	AllowCredentials: true,
//	 }))

	router.GET("/", welcome)
	router.POST("/api/login", authenticate)

	vmRouter := router.Group("/api/vms")
	vmRouter.GET("", getVMs)
	vmRouter.POST("", createVM)

	vmRouter.GET("/:id", getVM)
	vmRouter.DELETE("/:id", deleteVM)

	vmRouter.GET("/:id/monitor", monitorVM)

	vmRouter.GET("/:id/snapshots", snapshotlistVM)
	vmRouter.POST("/:id/snapshots", snapshotcreateVM)

	vmRouter.DELETE("/:id/snapshots/:sid", snapshotdeleteVM)
	vmRouter.POST("/:id/snapshots/:sid", snapshotrestoreVM)

	vmRouter.POST("/:id/actions", actionVM)

	vmRouter.POST("/:id/vnc", startVNC)

	router.POST("/:userid/resetpassword", passwordChange)	

	templateRouter := router.Group("/api/templates")
	templateRouter.GET("", getTemplates)
	templateRouter.GET("/:id", getTemplate)
	router.GET("/api/logout", removecookie)

	if *isProd {
		router.Logger.Fatal(router.Start(":80"))
	} else {
		router.Logger.Fatal(router.Start(":9090"))
	}
}
