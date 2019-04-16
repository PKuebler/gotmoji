package models

type Commit struct {
	Selected string `survey:"gitmoji"`
	Title    string `survey:"title"`
	Message  string `survey:"message"`
	Issue    string `survey:"issue"`
	Gitmoji  *Gitmoji
}
