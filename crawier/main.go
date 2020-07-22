package main

import (
	"crawier/engine"
	"crawier/zhenai/parser"
)
func main () {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}


