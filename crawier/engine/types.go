package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items   []interface{}
}

func NiParser([]byte) ParseResult {
	return ParseResult{}
}