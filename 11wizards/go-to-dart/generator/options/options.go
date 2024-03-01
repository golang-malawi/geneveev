package options

type Mode string

const (
	JSON      Mode = "json"
	Firestore Mode = "firestore"
)

type Options struct {
	Input   string
	Output  string
	Imports []string
	Mode    Mode
}
