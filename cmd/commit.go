package cmd

import (
	"bytes"
	"fmt"
	. "github.com/logrusorgru/aurora"
	"gopkg.in/AlecAivazis/survey.v1"
	"io"
	"log"
	"os"
	"os/exec"
)

func Commit() {
	config := LoadConfig()

	gitmoji := []string{}
	emojiMap := map[string]string{}

	for _, emoji := range config.Gitmojis {
		option := fmt.Sprintf("%s - %s (%s)", emoji.Emoji, emoji.Description, Magenta(emoji.Code))
		gitmoji = append(gitmoji, option)
		emojiMap[option] = emoji.Code
	}

	// the questions to ask
	var qs = []*survey.Question{
		{
			Name: "gitmoji",
			Prompt: &survey.Select{
				Message: "Choose a gitmoji",
				Options: gitmoji,
				Default: gitmoji[0],
			},
			Validate: survey.Required,
		},
		{
			Name:     "title",
			Prompt:   &survey.Input{Message: "Enter the commit title"},
			Validate: survey.Required,
		},
		{
			Name:   "message",
			Prompt: &survey.Input{Message: "Enter the commit message"},
		},
		{
			Name:   "issue",
			Prompt: &survey.Input{Message: "Issue / PR reference #"},
		},
		{
			Name:   "signed",
			Prompt: &survey.Confirm{Message: "Signed commit"},
		},
	}

	// the answers will be written to this struct
	answers := struct {
		Gitmoji string
		Title   string
		Message string
		Issue   string
		Signed  bool
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	issue := ""
	if len(answers.Issue) > 0 {
		issue = fmt.Sprintf(" (%s)", answers.Issue)
	}

	title := fmt.Sprintf("%s %s", emojiMap[answers.Gitmoji], answers.Title)
	body := fmt.Sprintf("%s%s", answers.Message, issue)

	signed := ""
	if answers.Signed {
		signed = "-S"
	}

	commit := []string{"commit", signed, fmt.Sprintf("-m \"%s\"", title), fmt.Sprintf("-m \"%s\"", body)}
	fmt.Sprintf("commit %s -m \"%s\" -m \"%s\"", signed, title, body)

	Run(commit)
}

func Run(commit []string) {
	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command("git", commit...)

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}
