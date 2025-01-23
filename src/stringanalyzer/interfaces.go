package stringanalyzer

type ICanGetJSON interface {
	GetJson() (*[]byte, error)
}

type IAnalyzedString interface {
	ICanGetJSON
	GetSize() int
	GetSymbolMap() map[rune]uint
}
