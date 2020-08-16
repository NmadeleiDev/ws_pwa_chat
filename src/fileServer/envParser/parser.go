package envParser

import (
	"fileServer/envParser/impl"
	"sync"

)

var syncOnce sync.Once
var parser Parser

type Parser interface {
	GetS3BucketPath() string
	IsDevMode() bool
	GetServerPort() string
}

func GetParser() Parser {
	syncOnce.Do(func() {
		parser = &impl.Parser{}
	})
	return parser
}
