package parser

import (
	"golang-demo/reptile/engine"
)

func NilParserFunc([]byte) engine.ParseResult {
	return engine.ParseResult{}
}
