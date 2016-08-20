package main

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl"
)

type fileConfig struct {
	LogPriority   string `hcl:"log_priority"`
	StateFilename string `hcl:"state_file"`
	BufferSize    int    `hcl:"buffer_size"`
}

type Config struct {
	LogPriority    Priority
	StateFilename  string
	BufferSize     int
}

func getLogLevel(priority string) (Priority, error) {

	logLevels := map[Priority][]string{
		EMERGENCY: {"0", "emerg"},
		ALERT: {"1", "alert"},
		CRITICAL: {"2", "crit"},
		ERROR: {"3", "err"},
		WARNING: {"4", "warning"},
		NOTICE: {"5", "notice"},
		INFO: {"6", "info"},
		DEBUG: {"7", "debug"},
	}

	for i, s := range logLevels {
	    if s[0] == priority || s[1] == priority {
	        return i, nil
	    }
	}

	return DEBUG, fmt.Errorf("'%s' is unsupported log priority", priority)
}

func LoadConfig(filename string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var fConfig fileConfig
	err = hcl.Decode(&fConfig, string(configBytes))
	if err != nil {
		return nil, err
	}

	if fConfig.StateFilename == "" {
		return nil, fmt.Errorf("state_file is required")
	}


	config := &Config{}

	if fConfig.LogPriority == "" {
		// Log everything
		config.LogPriority = DEBUG
	} else {
		config.LogPriority, err = getLogLevel(fConfig.LogPriority)
		if err != nil {
			return nil, fmt.Errorf("The provided log filtering '%s' is unsupported by systemd!", fConfig.LogPriority)
		}
	}

	if fConfig.BufferSize != 0 {
		config.BufferSize = fConfig.BufferSize
	} else {
		config.BufferSize = 100
	}


	return config, nil
}

