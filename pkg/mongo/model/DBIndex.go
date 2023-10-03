package mongoModel

type DBIndexType int32

const (
	DBIndexAsc  DBIndexType = 1
	DBIndexDesc DBIndexType = -1
)

type DBIndex struct {
	Collection Collection
	Name       string
	Keys       []string
	Type       DBIndexType
	Uniq       bool
}

type DBTextIndex struct {
	Collection string
	Name       string
	Keys       []string
}
