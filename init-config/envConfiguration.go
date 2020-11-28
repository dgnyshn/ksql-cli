package init_config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	DefaultEnvironment    = "dev"
	DefaultAuthentication = ""
	DefaultKsqlServerURL  = "http://localhost:8088"
)

const FileName = "ksqlCLIEnvConfigs.txt"

/*
	Config Types can be authentication and url
*/
type envConfigDetails struct {
	ConfigType string `json:"ConfigType"`
	Detail     string `json:"Detail"`
}

type configurationDetails struct {
	Env              string              `json:"Env"`
	Active           bool                `json:"Active"`
	EnvConfigDetails [2]envConfigDetails `json:"EnvConfigDetails"`
}

type configurations struct {
	ConfigurationDetails [3]configurationDetails `json:"ConfigurationDetails"`
}

func init() {
	if _, err := os.Stat(FileName); err == nil {
		initializeSelectedConfigs(err)
	} else {
		initializeDefaultConfigs(err)
	}
}

func initializeSelectedConfigs(err error) {
	file, err := ioutil.ReadFile(FileName)
	if err != nil {
		panic(err)
	}

	var configs configurations
	json.Unmarshal(file, &configs)

	mapInitConfigs(configs)
}

func mapInitConfigs(configs configurations) {
	for _, configDetail := range configs.ConfigurationDetails {
		if configDetail.Active == true {
			for _, envConfigDetail := range configDetail.EnvConfigDetails {
				if envConfigDetail.ConfigType == "authentication" {
					DefaultAuthentication = envConfigDetail.Detail
				} else if envConfigDetail.ConfigType == "url" {
					DefaultKsqlServerURL = envConfigDetail.Detail
				}
			}
		}
	}

}

func initializeDefaultConfigs(err error) {
	if os.IsNotExist(err) {
		create, err := os.Create(FileName)
		if err != nil {
			panic(err)
		}

		defer create.Close()

		c := defaultConfigs()
		jsonConfig, _ := json.Marshal(c)

		create.Write(jsonConfig)

	}
}

func defaultConfigs() configurations {
	return configurations{
		ConfigurationDetails: [3]configurationDetails{
			{
				Env:    "dev",
				Active: true,
				EnvConfigDetails: [2]envConfigDetails{
					{
						ConfigType: "authentication",
						Detail:     "",
					},
					{
						ConfigType: "url",
						Detail:     "http://localhost:8088",
					},
				},
			},
			{
				Env:    "stage",
				Active: false,
				EnvConfigDetails: [2]envConfigDetails{
					{
						ConfigType: "authentication",
						Detail:     "",
					},
					{
						ConfigType: "url",
						Detail:     "",
					},
				},
			},

			{
				Env:    "prod",
				Active: false,
				EnvConfigDetails: [2]envConfigDetails{
					{
						ConfigType: "authentication",
						Detail:     "",
					},
					{
						ConfigType: "url",
						Detail:     "",
					},
				},
			},
		},
	}
}

func Save(in string) {
	file, err := ioutil.ReadFile(FileName)

	if err != nil {
		return
	}

	var configs configurations
	commands := strings.Fields(strings.ToLower(in))
	if len(file) > 0 {
		err := json.Unmarshal(file, &configs)
		if err != nil {
			fmt.Println("Exception occurred while reading file")
		}

		updateEnvironmentConfigs(&configs, commands)

	} else {
		fmt.Println("Exception occurred while preparing default file")
	}

	preparedFile, _ := json.MarshalIndent(configs, "", " ")
	_ = ioutil.WriteFile(FileName, preparedFile, 0644)
}

func ActivateEnv(in string) {
	file, err := ioutil.ReadFile(FileName)

	if err != nil {
		return
	}

	var configs configurations
	commands := strings.Fields(strings.ToLower(in))

	if len(file) > 0 {
		err := json.Unmarshal(file, &configs)
		if err != nil {
			fmt.Println("Exception occurred while reading file")
		}

		activate(&configs, commands)
		mapInitConfigs(configs)
	} else {
		fmt.Println("Exception occurred while preparing default file")
	}

	preparedFile, _ := json.MarshalIndent(configs, "", " ")
	_ = ioutil.WriteFile(FileName, preparedFile, 0644)

}

func activate(configs *configurations, commands []string) {
	for i := range configs.ConfigurationDetails {
		if configs.ConfigurationDetails[i].Active {
			configs.ConfigurationDetails[i].Active = false
			break
		}
	}

	for i := range configs.ConfigurationDetails {
		if configs.ConfigurationDetails[i].Env == commands[len(commands)-1] {
			configs.ConfigurationDetails[i].Active = true
		}
	}
}

func updateEnvironmentConfigs(configs *configurations, commands []string) bool {
	for i := range configs.ConfigurationDetails {
		details := configs.ConfigurationDetails[i]
		if details.Env == getEnvironment(commands) {
			for j := range details.EnvConfigDetails {
				configDetails := details.EnvConfigDetails[j]
				if configDetails.ConfigType == commands[len(commands)-2] {
					configs.ConfigurationDetails[i].EnvConfigDetails[j].Detail = commands[len(commands)-1]
					return true
				}
			}
		}

	}
	return false
}

func getEnvironment(commands []string) string {
	for _, env := range commands {
		if isEnvExists(env) {
			return env
		}
	}
	panic("Environment could not find please enter an environment")
}

func isEnvExists(env string) bool {
	return strings.Compare(env, "dev") == 0 ||
		strings.Compare(env, "stage") == 0 ||
		strings.Compare(env, "prod") == 0
}
