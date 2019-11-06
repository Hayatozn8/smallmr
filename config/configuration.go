package config

import (
	"strings"
)

type Configuration struct {
	int64Conf map[string]int64
	int32Conf map[string]int32
	uint32Conf map[string]uint32

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

		uint32Conf:map[string]uint32{
			MAX_TASK_COUNT:10, // TODO
		},

		stringConf:map[string]string{
			FILE_DELIMITER:"",
			INPUT_PATHS:"",
		},
	}
}

func (conf *Configuration) SetInputPaths(paths ...string){
	if len(paths) == 0 {
		return
	}

	conf.SetString(INPUT_PATHS, strings.Join(paths, INPUT_PATHS_JOIN_SEP))
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

func (conf *Configuration) GetUint32(key string) uint32 {
	return conf.uint32Conf[key]
}

func (conf *Configuration) SetUint32(key string, value uint32){
	conf.uint32Conf[key] = value
}

func (conf *Configuration) GetString(key string) string {
	return conf.stringConf[key]
}

func (conf *Configuration) SetString(key string, value string){
	conf.stringConf[key] = value
}
