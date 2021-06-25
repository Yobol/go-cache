package cache

type Stat struct {
	Pairs     int64 `json:"pairs"`      // the number of k-v pairs currently stored by the cache server
	KeySize   int64 `json:"key_size"`   // the size of all keys currently stored by the cache server
	ValueSize int64 `json:"value_size"` // the size of all values currently stored by the cache server
}

func (s *Stat) add(key string, value []byte) {
	s.Pairs++
	s.KeySize += int64(len(key))
	s.ValueSize += int64(len(value))
}

func (s *Stat) del(key string, value []byte) {
	s.Pairs--
	s.KeySize -= int64(len(key))
	s.ValueSize -= int64(len(value))
}
