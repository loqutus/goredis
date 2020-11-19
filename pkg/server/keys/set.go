package keys

func Set(key string, value string) error {
	lock.Lock()
	defer lock.Unlock()
	Dict[key] = value
	return nil
}
