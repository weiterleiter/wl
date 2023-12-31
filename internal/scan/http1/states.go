package http1

type parserState int

const (
	eRequestLine parserState = iota
	eHeaderKey
	eHeaderKeyCR
	eHostValue
	eContentLengthValue
	eContentLengthValueCR
	eOtherHeaderValue
	eBody
)

type chunkedState int

const (
	eChunkLength chunkedState = iota
	eChunkBody
	eLastChunk
)
