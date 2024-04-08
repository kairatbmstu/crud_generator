package parser

const (
	Keyword_Entity       = "entity"
	Keyword_Relationship = "relationship"
	Keyword_Paginate     = "paginate"
)

type DataType int

const (
	DataTypeUndefined   DataType = 0
	DataTypeText        DataType = 1
	DataTypePunctuation DataType = 2
	DataTypeWhiteSpace  DataType = 3
)
