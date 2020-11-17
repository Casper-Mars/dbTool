package pojo

type TableInfo struct {
	TableName string
	Comment   string
	Cols      []ColInfo
}

type ColInfo struct {
	ColName   string
	Comment   string
	ColType   string
	IsEmpty   bool
	Len       int
	IsPrimary bool
}
