package keys

import "errors"

func Get(key string) (string, error) {
	lock.RLock()
	defer lock.RUnlock()
	if val, ok := Dict[key]; ok {
		return val, nil
	} else {
		return "", errors.New("key not found")
	}
}
