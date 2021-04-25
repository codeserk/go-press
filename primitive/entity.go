package primitive

type Type int

const (
	Text Type = iota
)

type Entity struct {
	ID    string `json:"id" example:"8d6dde41-d23a-49a5-a9d7-6cbba52ec7fc"`
	Name  string `json:"name" example:"Ryuk"`
	Color string `json:"color" example:"Black" description:"Cat color"`
}
