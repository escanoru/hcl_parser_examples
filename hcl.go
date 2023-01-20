package main

import (
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type LogLevl struct {
	LogLevel string `hcl:"log_level"`
}

type Config struct {
	IOMode  string        `hcl:"io_mode"`
	Service ServiceConfig `hcl:"service,block"`
}

type ServiceConfig struct {
	Protocol   string          `hcl:"protocol,label"`
	Type       string          `hcl:"type,label"`
	ListenAddr string          `hcl:"listen_addr"`
	Processes  []ProcessConfig `hcl:"process,block"`
}

type ProcessConfig struct {
	Type    string   `hcl:"type,label"`
	Command []string `hcl:"command"`
}

func main() {
	var config Config
	var config2 LogLevl
	err := hclsimple.DecodeFile("config.hcl", nil, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	log.Printf("Configuration is %#v", config)

	err = hclsimple.DecodeFile("config2.hcl", nil, &config2)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	log.Printf("Configuration file is %#v", config2)
}
