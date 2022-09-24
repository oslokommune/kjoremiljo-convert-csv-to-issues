package convert

type Issue struct {
	Label string
	Title string
	Body  string

	moveToGithub bool
}
