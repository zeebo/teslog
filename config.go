package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/zeebo/teslog/teslib"
)

type Config struct {
	Creds   teslib.Creds
	Vehicle int64
}

var (
	teslogDir  = os.ExpandEnv("$HOME/.teslog")
	configFile = os.ExpandEnv("$HOME/.teslog/config")
)

func promptString(prompt string) (out string, err error) {
	fmt.Print(prompt + ": ")
	_, err = fmt.Scanln(&out)
	return out, err
}

func promptCreds() (creds teslib.Creds, err error) {
	creds.Email, err = promptString("Email")
	if err != nil {
		return creds, err
	}

	creds.Password, err = promptString("Password")
	if err != nil {
		return creds, err
	}

	return creds, nil
}

func loadConfig() (config Config, ok bool) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return config, false
	}
	err = json.Unmarshal(data, &config)
	return config, err == nil
}

func saveConfig(config Config) {
	if data, err := jsonEncode(config); err == nil {
		os.MkdirAll(teslogDir, 0755)
		ioutil.WriteFile(configFile, data, 0600)
	}
}
