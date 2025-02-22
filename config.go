package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"regexp"
)

type LogReader struct {
	Command          string `yaml:"command"`
	SingleLogCommand string `yaml:"single_log_command"`
}

type Test struct {
	MatchPathRegex          regexp.Regexp `yaml:"match_path_regex"`
	PreTestCommand          string        `yaml:"pre_test_command"`
	TestAllCommand          string        `yaml:"test_all_command"`
	TestCommand             string        `yaml:"test_command"`
	TestCommandSearchPrefix string        `yaml:"test_command_search_prefix"`
	FailedLogListCommand    string        `yaml:"failed_log_list_command"`
}

type Tester struct {
	Tests     []Test    `yaml:"tests"`
	LogReader LogReader `yaml:"log_reader"`
}

type Config struct {
	Tester Tester `yaml:"tester"`
}

func ReadConfig(filename string) Config {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	config := Config{}
	if err := yaml.Unmarshal(raw, &config); err != nil {
		panic(err)
	}

	return config
}
