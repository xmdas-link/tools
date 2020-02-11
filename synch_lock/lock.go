package synch_lock

type Lock struct {
	Key string
}

func (l *Lock) IsFree() bool {
	return l.Key == ""
}

func (l *Lock) IsBlock(key string) bool {
	return l.Key == key
}

func (l *Lock) Release(key string) {
	if l.Key == key {
		//fmt.Println("[Release]" + key)
		l.Key = ""
	}
}

func (l *Lock) Take(key string) {
	//fmt.Println("[Take]" + key)
	l.Key = key
}
