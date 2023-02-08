package domain

import (
	"handshake/domain/role"
	topic2 "handshake/domain/topic"
	user2 "handshake/domain/user"
)

func Init() {
	role.ListExample.Init()
	user2.ListExample.Init()
	topic2.ListExample.Init()
}
