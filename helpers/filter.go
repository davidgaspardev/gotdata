package helpers

type Where struct {
	Attribute string
	Operator  string
	Value     interface{}
}

type Filter struct {
	Wheres []*Where
	Page   uint
	Orders []string
}
