package upload

type File struct {
	Type      FileType
	Suffix    []string
	MaxSize   int
	UrlPrefix string
	LocalPath string
}

func NewFile(filetype FileType, suffix []string, maxSize int, urlPrefix string, localpath string) *File {
	return &File{Type: filetype, Suffix: suffix, MaxSize: maxSize, UrlPrefix: urlPrefix, LocalPath: localpath}
}

func (f *File) GetType() FileType {
	return f.Type
}

func (f *File) GetSuffix() []string {
	return f.Suffix
}

func (f *File) GetMaxSize() int {
	return f.MaxSize
}

func (f *File) GetUrlPrefix() string {
	return f.UrlPrefix
}

func (f *File) GetPath() string {
	return f.LocalPath
}
