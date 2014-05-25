package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"encoding/json"
	"github.com/albrow/zoom"
	"github.com/dchest/authcookie"
	"github.com/revel/revel"
	"time"
	"user-auth/app/models"
)

var appSecret []byte = []byte("21d6306e28f0d7b1bbb38cad07ad378356e68931d08a95c4")

type Users struct {
	*revel.Controller
	JsonController
}

type AuthData struct {
	Email          string
	HashedPassword string
}

func (c Users) Authenticate(email string, password string) revel.Result {

	return c.Render()
}

func (c Users) Login(email string, password string) revel.Result {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return c.RenderJsonError(500, err)
	}
	q := zoom.NewQuery("User").Filter("HashedPassword =", string(hash)).Filter("Email =", email)
	num, err := q.Count()
	if err != nil {
		return c.RenderJsonError(500, err)
	}
	if num == 0 {
		return c.RenderJsonError(400, err)
	}
	data := AuthData{Email: email, HashedPassword: string(hash)}
	dataJson, err := json.Marshal(data)
	if err != nil {
		return c.RenderJsonError(500, err)
	}
	cookie := authcookie.NewSinceNow(string(dataJson), 24*time.Hour, appSecret)
	c.Session["authentication"] = cookie
	return c.RenderJsonOk()
}

func (c Users) Create(email string, password string) revel.Result {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return c.RenderJsonError(500, err)
	}
	u := &models.User{Email: email, HashedPassword: string(hash)}
	if err := zoom.Save(u); err != nil {
		return c.RenderJsonError(500, err)
	}
	return c.RenderJson(u)
}
