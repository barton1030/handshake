package domain

import (
	"handshake/domain/role"
	topic2 "handshake/domain/topic"
	user2 "handshake/domain/user"
)

func Init() {
	role.List.Init()
	user2.List.Init()
	topic2.List.Init()
}
