package controls

import (
	"bufio"
	"os"
	"pathbot/locations"
	"strings"
)

type KeyboardInput struct {
	Prompt string
	Reader *bufio.Reader
}

var DefaultInput = KeyboardInput{
	Prompt:"What direction you will go?",
	Reader: bufio.NewReader(os.Stdin),
}

func (k *KeyboardInput) UserInput() (dir *locations.PathBotDirection, err error) {
	input, err := k.Reader.ReadString('\n')
	if err != nil { return nil, err }
	dir = &locations.PathBotDirection{Direction: strings.ToUpper(input[0:1])}
	return dir, nil
}