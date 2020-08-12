package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string
	Password string
}

type SignUpStruct struct {
	Username        string
	Password        string
	PasswordConfirm string
}

func (x SignUpStruct) toUser() User {
	return User{x.Username, x.Password}
}

type SignInStruct struct {
	Username string
	Password string
}

func (x SignInStruct) toUser() User {
	return User{x.Username, x.Password}
}

var userDatabase map[string]User

func GetUser(username string) (User, bool) {
	val, ok := userDatabase[username]
	return val, ok
}

func SignUp(user *SignUpStruct) (*User, bool) {
	if user.Password != user.PasswordConfirm {
		return nil, false
	}

	_, exists := GetUser(user.Username)
	if exists {
		return nil, false
	}

	u := user.toUser()
	userDatabase[u.Username] = u

	return &u, true
}

func SignIn(user *SignInStruct) (*User, bool) {
	u, exists := GetUser(user.Username)
	if !exists {
		return nil, false
	}

	if user.Password != u.Password {
		return nil, false
	}

	return &u, true
}

func IndexGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		name, err := c.Cookie("username")
		username := "none"
		if err == nil {
			user, exists := GetUser(name)
			if exists {
				username = "@" + user.Username
			}
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": username,
		})
	}
}

func DevelopGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		name, err := c.Cookie("username")
		username := "none"
		if err == nil {
			user, exists := GetUser(name)
			if exists {
				username = "@" + user.Username
			}
		}

		c.HTML(http.StatusOK, "develop.html", gin.H{
			"username": username,
		})
	}
}

func SignUpPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.ParseForm()
		user := SignUpStruct{
			c.Request.Form["username"][0],
			c.Request.Form["password"][0],
			c.Request.Form["password_confirm"][0],
		}

		u, ok := SignUp(&user)
		if ok {
			// c.SetCookie(name, val, maxAge, path, domain, secure, httpOnly)
			c.SetCookie("username", u.Username, 300, "/", "", false, true)
		}

		c.Redirect(http.StatusSeeOther, "/")
	}
}

func SignInPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.ParseForm()
		user := SignInStruct{
			c.Request.Form["username"][0],
			c.Request.Form["password"][0],
		}

		u, ok := SignIn(&user)
		if ok {
			// c.SetCookie(name, val, maxAge, path, domain, secure, httpOnly)
			c.SetCookie("username", u.Username, 300, "/", "", false, true)
		}

		c.Redirect(http.StatusSeeOther, "/")
	}
}

func SignOutPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("username", "", -1, "/", "", false, true)
		c.Redirect(http.StatusSeeOther, "/")
	}
}

func main() {
	userDatabase = make(map[string]User)

	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	r.GET("/", IndexGet())
	r.GET("/develop", DevelopGet())

	r.POST("/signup", SignUpPost())
	r.POST("/signin", SignInPost())
	r.POST("/signout", SignOutPost())

	r.Run(":8080")
}
