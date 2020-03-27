package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/OpenNebula/one/src/oca/go/src/goca"
	"github.com/labstack/echo/v4"
)

type (
	credentials struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
		Remember string `json:"remember_me" form:"remember_me"`
	}
)

var (
	// OpenNebula Panel URL
	xonglURL = "https://cloud.xongl.com"
	// OpenNebula API URL
	apiURL = "http://cloud.xongl.com:2633/RPC2"
	// user: one object
	users  = make(map[string]*goca.Controller)
	config = make(map[string]goca.OneConfig)
)

func authenticate(c echo.Context) error {
	user := new(credentials)
	if err := c.Bind(user); err == nil {
		if user.Username != "" && user.Password != "" {
			conf := goca.NewConfig(user.Username, user.Password, apiURL)
			client := goca.NewDefaultClient(conf)
			controller := goca.NewController(client)
			_, err := controller.Templates().Info()
			if err == nil {
				expirationTime := time.Now().Add(720 * time.Hour)
				tokenString, err := generateToken(user.Username, getSecretKey(), expirationTime.Unix())
				if err == nil {
					c = setCookie(c, tokenString, expirationTime, user.Remember)
					users[user.Username] = controller
					config[user.Username] = conf
					return c.JSON(http.StatusOK, map[string]string{
						"token": tokenString,
					})
				}
				return echo.ErrInternalServerError
			}
			return c.JSON(http.StatusUnauthorized, errorMsg(err))
		}
		return c.JSON(http.StatusBadRequest, errorJSON{"Invalid Credentials!"})
	}
	return c.JSON(http.StatusBadRequest, errorJSON{"Invalid Data"})
}

func setCookie(c echo.Context, tokenString string, expirationTime time.Time, remember string) echo.Context {
	hps := strings.Split(tokenString, ".")
	if remember != "" {
		c.SetCookie(&http.Cookie{
			Name:     "xonglTokenHAP",
			Value:    hps[0] + "." + hps[1],
			SameSite: http.SameSiteStrictMode,
			Secure:   *isProd,
			Expires:  expirationTime,
		})
		c.SetCookie(&http.Cookie{
			Name:     "xonglTokenSIG",
			Value:    hps[2],
			HttpOnly: true,
			Secure:   *isProd,
			SameSite: http.SameSiteStrictMode,
			Expires:  expirationTime,
		})
	} else {
		c.SetCookie(&http.Cookie{
			Name:     "xonglTokenHAP",
			Value:    hps[0] + "." + hps[1],
			SameSite: http.SameSiteStrictMode,
			Secure:   *isProd,
		})
		c.SetCookie(&http.Cookie{
			Name:     "xonglTokenSIG",
			Value:    hps[2],
			HttpOnly: true,
			Secure:   *isProd,
			SameSite: http.SameSiteStrictMode,
		})
	}
	return c
}
func removecookie(c echo.Context) error{

	c.SetCookie(&http.Cookie{
		Name:     "xonglTokenHAP",
		Value:    "",
		SameSite: http.SameSiteStrictMode,
		Secure:   *isProd,
		Expires: time.Now().Add(-1 * time.Hour),
	})
	c.SetCookie(&http.Cookie{
		Name:     "xonglTokenSIG",
		Value:    "",
		HttpOnly: true,
		Secure:   *isProd,
		SameSite: http.SameSiteStrictMode,
	})
	return c.JSON(http.StatusOK, okMsg("Logged Out"))

}
