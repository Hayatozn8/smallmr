package config

type Configuration struct {
	int64Conf map[string]int64
}

func (conf *Configuration) GetInt64(key string) int64 {
	return conf.int64Conf[key]
}
