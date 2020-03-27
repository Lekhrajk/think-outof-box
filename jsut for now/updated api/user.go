package main


import (
        "net/http"
        "strconv"
        "github.com/OpenNebula/one/src/oca/go/src/goca"
        "github.com/labstack/echo/v4"
)


type (
	// User
	user  struct {
		Password string `json:"password"`
	}
)

func passwordChange(c echo.Context) error {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err == nil {
		config := new(user)
                one := c.Get("one").(*goca.Controller)
                userid = -1 
		if err := c.Bind(config); 
			err == nil {
                                err = one.User(userid).Passwd(config.Password)
                        if err == nil {
                                return c.JSON(http.StatusCreated, okMsg("Password Changed"))
                        }
                        return c.JSON(http.StatusBadRequest, errorMsg(err))
                }
                return c.JSON(http.StatusBadRequest, errorMsg(err))
        }
        return c.JSON(http.StatusBadRequest, errorJSON{"Invalid Password"})
}

