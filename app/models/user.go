package models

import (
	"github.com/albrow/zoom"
)

type User struct {
	Email          string `zoom:"index"`
	HashedPassword string `zoom:"index"`
	zoom.DefaultData
}
