package main

import (
	"encoding/json"

	"goim/libs/define"
)

// developer could implement "Auth" interface for decide how get userId, or roomId
type Auther interface {
	Auth(token string) (userId int64, roomId int32)
}

type DefaultAuther struct {
}

func NewDefaultAuther() *DefaultAuther {
	return &DefaultAuther{}
}

// func (a *DefaultAuther) Auth(token string) (userId int64, roomId int32) {
// 	var err error
// 	if userId, err = strconv.ParseInt(token, 10, 64); err != nil {
// 		userId = 0
// 		roomId = define.NoRoom
// 	} else {
// 		roomId = 1 // only for debug
// 	}
// 	return
// }

func (a *DefaultAuther) Auth(token string) (userId int64, roomId int32) {
	auth := &Auth{}
	json.Unmarshal([]byte(token), &auth)

	roomId = auth.Roomid
	userId = auth.Uid
	if roomId < 0 {
		roomId = define.NoRoom
	}
	return
}

type Auth struct {
	Roomid int32 `json:"roomid"`
	Uid    int64 `json:"uid"`
}
