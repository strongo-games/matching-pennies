package mpmodels

import (
	"github.com/strongo/db"
	"github.com/strongo/app"
	"github.com/strongo/bots-framework/core"
)

const UserKind = "User"

type User struct {
	db.IntegerID
	*UserEntity
}

type UserEntity = strongo.AppUserBase


var _ bots.BotAppUser = (*User)(nil)