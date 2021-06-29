package tag

type TagType = string

const (
	IPAnyCastTag TagType = "IP AnyCast"
	ServerTag    TagType = "Server Tag"
)

type Tag struct {
	ID    string
	Title string
	TagType
}
