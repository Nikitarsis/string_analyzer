package string_analyzer

/*
Interface allow json serialization
*/
type ICanGetJSON interface {
	GetJson() (*[]byte, error)
}

/*
Analyzed String interface
*/
type IAnalyzedString interface {
	ICanGetJSON
	GetSize() int
	GetSymbolMap() map[string]uint
	GetFlagMap() map[string]bool
}
