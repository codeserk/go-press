package primitive

type Type string

const (
	Text Type = "text"
)

type Entity struct {
	ID   string `json:"id"`
	Name string `json:"name" example:"Ryuk"`
}
