package main

import (
	"./commands"
	config "./initCfg"
	serverStatus "./ksql-server/server-detail"
	streams "./ksql-server/statements"
	"fmt"
	"github.com/c-bata/go-prompt"
	"os"
	"os/exec"
	"strings"
)

func completer(t prompt.Document) []prompt.Suggest {
	c := commands.New()
	text := t.CurrentLine()

	if len(text) < 1 {
		return []prompt.Suggest{}
	}

	if len(GetFormattedCommand(text)) > 2 {
		key := GetFormattedCommand(t.GetWordBeforeCursorWithSpace())
		return c.GetSubSuggestions(key)
	}

	return c.GetSuggestions()
}

func executor(in string) {
	command := in
	if command == "exit" {
		os.Exit(0)
	} else if command == "clear" {
		executeOsCommand(in)
		return
	}

	splitCommand := strings.Fields(command)

	switch strings.ToLower(splitCommand[0]) {
	case "list":
		lastCommand := splitCommand[1]
		switch lastCommand {
		case "stream", "streams", "streams;":
			streams.ListStreams()
			break
		case "tables", "tables;":
			streams.ListTables()
			break
		case "queries", "queries;":
			streams.ListQueries()
			break
		case "topics", "topics;":
			streams.ListTopics()
			break
		case "prop", "property", "properties", "properties;":
			streams.ListProperties()
			break
		default:
			fmt.Println("Command could not executable")
			break
		}
	case "info":
		serverStatus.Info()
		break
	case "healthcheck":
		serverStatus.HealthCheck()
		break
	case "create", "drop", "terminate", "insert":
		streams.ExecuteCrudStatement(in)
		break

	case "describe":
		streams.ExecDescribe(in)
		break
	case "select":
		streams.Execute(in)
		break
	case "print":
		streams.ExecutePrintTopic(in)
		break
	case "save":
		if isEnvUpdate(splitCommand) {
			config.Save(in)
		}
		break
	case "use":
		config.ActivateEnv(in)
		break
	default:
		fmt.Println("Command could not executable")
		break
	}
}

func isEnvUpdate(commands []string) bool {
	if len(commands) < 2 {
		return false
	}

	return commands[len(commands)-2] == "authentication" ||
		commands[len(commands)-2] == "url"
}

func main() {
	command := prompt.New(executor, completer,
		prompt.OptionTitle("KsqlServer Prompt"),
		prompt.OptionPrefix("ksql-cli-> "),
		prompt.OptionPrefixBackgroundColor(prompt.DefaultColor),
		prompt.OptionSuggestionTextColor(prompt.Black),
	)
	command.Run()
}

func GetFormattedCommand(command string) string {
	return strings.ToLower(strings.TrimSpace(command))
}

func executeOsCommand(in string) {
	command := exec.Command(in)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
}
