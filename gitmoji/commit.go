package gitmoji

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/pkuebler/gotmoji/models"
)

func Commit(answers models.Commit, useUTF8 bool) {
	emoji := answers.Gitmoji.Code

	if useUTF8 {
		emoji = answers.Gitmoji.Emoji
	}

	title := fmt.Sprintf("-m \"%s %s %s\"", emoji, answers.Title, answers.Issue)
	message := ""

	commit := []string{"commit"}
	commit = append(commit, title)

	if len(answers.Message) > 0 {
		message = fmt.Sprintf("-m \"%s\"", answers.Message)
		commit = append(commit, message)
	}

	fmt.Printf("git commit %s %s\n", title, message)

	RunCommit(commit)
}

func RunCommit(commit []string) {
	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command("git", commit...)

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Start()
	if err != nil {
		fmt.Println("cmd.Start() failed with '%s'\n", err)
		os.Exit(1)
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
