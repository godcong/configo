package configo

import "errors"

//error
var (
	ErrorConfigUnknown          = errors.New("unknown")
	ErrorConfigNotFound         = errors.New("config not found")
	ErrorConfigCannotOpen       = errors.New("config cannot open")
	ErrorConfigGetProperty      = errors.New("property not found")
	ErrorConfigGetPropertyType  = errors.New("property type not found")
	ErrorConfigGetPropertyValue = errors.New("property value not found")
	ErrorSplitNoData            = errors.New("split no data")
	ErrorGetNothing             = errors.New("get nothing")
	ErrorGetMapValueNotFound    = errors.New("map value not found ")
)
