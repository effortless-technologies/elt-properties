package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func login(c echo.Context) error {

	type User struct {
		Username		string			`json:"username"`
		Password		string			`json:"password"`
	}

	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	if u.Username == "ef" && u.Password == "1234" {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func accessible(c echo.Context) error {

	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {

	e := echo.New()
	e.Use(middleware.CORS())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", login)

	e.GET("/", accessible)

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":7000"))
}

