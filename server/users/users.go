package users

import "sync"

type User struct {
	Username       string
	HashedPassword string
}

var UserStore = make(map[string]*User)
var UserStoreLock sync.Mutex
