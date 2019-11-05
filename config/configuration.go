package config

type Configuration struct {
	int64Conf map[string]int64
	int32Conf map[string]int32
	
	stringConf map[string]string

}

func NewConfiguration() *Configuration {
	return &Configuration{
		int64Conf:map[string]int64{
			SPLIT_MAXSIZE: 33554432, // 32M=32 * 1024 * 1024 byte
			SPLIT_MINSIZE:1,
		},

		int32Conf:map[string]int32{
			MAX_LINE_LENGTH:33554432, // TODO
		},

		stringConf:map[string]string{
			FILE_DELIMITER:"",
		},
	}
}

func (conf *Configuration) GetInt64(key string) int64 {
	return conf.int64Conf[key]
}

func (conf *Configuration) SetInt64(key string, value int64){
	conf.int64Conf[key] = value
}

func (conf *Configuration) GetInt32(key string) int32 {
	return conf.int32Conf[key]
}

func (conf *Configuration) SetInt32(key string, value int32){
	conf.int32Conf[key] = value
}

func (conf *Configuration) GetString(key string) string {
	return conf.stringConf[key]
}

func (conf *Configuration) SetString(key string, value string){
	conf.stringConf[key] = value
}
