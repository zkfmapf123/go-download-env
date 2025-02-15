package interaction

import (
	"log"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	
	if err != nil {
		log.Fatalln(err)
	}
}

func SelectBox(msg string, list []string) string {
	prompt := &survey.Select{
		Message: msg,
		Options: list,
	}

	var answer string

	err := survey.AskOne(prompt, &answer, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return answer
}

func SelectMultipleBox(msg string, list []string) []string {
	prompt := &survey.MultiSelect{
		Message: msg,
		Options: list,
	}

	answer := []string{}
	err := survey.AskOne(prompt, &answer)
	if err != nil {
		log.Fatalln(err)
	}
	return answer
}

func InputText(msg string) string {
	prompt := &survey.Input{
		Message: msg,
	}

	var answer string
	err := survey.AskOne(prompt, &answer, nil)
	if err != nil {
		log.Fatalln(err)
	}
	
	return answer
}

func PressEnter(text string) string {
	prompt := &survey.Input{
		Message: text,
	}

	var answer string
	err := survey.AskOne(prompt, &answer, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return answer
}