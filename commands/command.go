package commands

import (
	"github.com/c-bata/go-prompt"
)

type Commands struct {
	Suggestions    []prompt.Suggest
	SubSuggestions map[string][]prompt.Suggest
}

func New() Commands {
	return Commands{
		Suggestions: []prompt.Suggest{
			{"info", "Get Information Details of KsqlDB server"},
			{"healthCheck", "Check the health of your ksqlDB server"},
			{"list", "List streams, topics , tables or queries"},
			{"exit", "Exit from the prompt"},
			{"save", "Save the environment details Password - User - Server URL. It makes configurable"},
			{"clear", "Clean the screen"},
			{"use", "Change the default usage env. prod - stage - dev"},
		},

		SubSuggestions: map[string][]prompt.Suggest{
			"list": {
				prompt.Suggest{
					Text: "topics", Description: "Get topic names",
				},
				prompt.Suggest{
					Text: "streams", Description: "Get streams names",
				},
				prompt.Suggest{
					Text: "tables", Description: "Get table names",
				},
				prompt.Suggest{
					Text: "queries", Description: "Get query names",
				},
				prompt.Suggest{
					Text: "properties", Description: "Get all properties",
				},
			},
			"save": {
				prompt.Suggest{
					Text: "env", Description: "Select Environment Environment and update",
				},
			},
			"env": {
				prompt.Suggest{
					Text: "dev", Description: "Local environment",
				},
				prompt.Suggest{
					Text: "stage", Description: "Stage environment",
				},
				prompt.Suggest{
					Text: "prod", Description: "Production environment",
				},
			},
			"dev": {
				prompt.Suggest{
					Text: "authentication", Description: "User name and password for authentication <username>:<password>",
				},
				prompt.Suggest{
					Text: "url", Description: "Localhost Ksql Server url",
				},
			},
			"stage": {
				prompt.Suggest{
					Text: "authentication", Description: "User name and password for authentication <username>:<password>",
				},
				prompt.Suggest{
					Text: "url", Description: "Stage Ksql Server url",
				},
			},
			"prod": {
				prompt.Suggest{
					Text: "authentication", Description: "User name and password for authentication <username>:<password>",
				},
				prompt.Suggest{
					Text: "url", Description: "Production Ksql Server url",
				},
			},

			"use": {
				prompt.Suggest{
					Text: "dev", Description: "Use dev environment to connect server",
				},
				prompt.Suggest{
					Text: "stage", Description: "Use stage environment to connect server",
				},
				prompt.Suggest{
					Text: "prod", Description: "Use production environment to connect server",
				},
			},
		},
	}
}

func (c *Commands) GetSuggestions() []prompt.Suggest {
	return c.Suggestions
}

func (c *Commands) GetSubSuggestions(key string) []prompt.Suggest {
	return c.SubSuggestions[key]
}
