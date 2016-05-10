package lib

import (
	"os/user"
)

func CurrentUserName() string {
	u, _ := user.Current()
	return u.Username
}
