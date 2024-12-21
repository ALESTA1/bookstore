package db

import (
	pd "bookstore/proto"
	"sync"
)

var Mu sync.Mutex
var BookMap = make(map[int32]*pd.Book)
