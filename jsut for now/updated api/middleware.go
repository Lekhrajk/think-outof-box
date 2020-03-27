package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	excludeRoutes = []string{"/", "/api/login"}
)

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Xongl Cloud\n")
}

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		for _, route := range excludeRoutes {
			if route == c.Request().URL.Path {
				return next(c)
			}
		}

		var tokenString string

		hap, haperr := c.Cookie("xonglTokenHAP")
		sig, sigerr := c.Cookie("xonglTokenSIG")
		if haperr == nil && sigerr == nil {
			tokenString = hap.Value + "." + sig.Value
			goto VERIFY
		}
		if authHeader := c.Request().Header.Get("Authorization"); authHeader != "" {
			tokenPart := strings.Split(authHeader, " ")
			if len(tokenPart) == 2 {
				tokenString = tokenPart[1]
			}
		}
	VERIFY:
		if tokenString != "" {
			config, ok, err := verifyToken(tokenString, jwtKey)
			if ok {
				if one, ok := users[config.Username]; ok {
					c.Set("one", one)
					return next(c)
				}
				return c.JSON(http.StatusUnauthorized, errorJSON{"Kindly get a new token"})
			}
			return c.JSON(http.StatusUnauthorized, errorJSON{err.Error()})
		}
		return c.JSON(http.StatusUnauthorized, errorJSON{"Missing cookie or authorization header!"})
	}
}
