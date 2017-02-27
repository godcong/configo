package config

import "errors"

//错误编号
var (
	ERROR_CONFIG_UNKNOWN          = errors.New("unknown")
	ERROR_CONFIG_NOT_FOUND        = errors.New("config not found")
	ERROR_CONFIG_CANNOT_OPEN      = errors.New("config cannot open")
	ERROR_SPLIT_NO_DATA           = errors.New("split no data")
	ERROR_GET_NOTHING             = errors.New("get nothing")
	ERROR_GET_MAP_VALUE_NOT_FOUND = errors.New("map value not found ")
)
