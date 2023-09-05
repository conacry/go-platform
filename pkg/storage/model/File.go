package storageModel

type File struct {
	Content  []byte
	Scope    string
	FilePath string
	MIME     string
}
