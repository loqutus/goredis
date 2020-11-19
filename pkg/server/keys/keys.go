package keys

import "sync"

var Dict map[string]string
var lock = sync.RWMutex{}
