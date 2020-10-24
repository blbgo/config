package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/blbgo/general"
)

// ErrSectionNotFound is the error returned when a config section is not found
var ErrSectionNotFound = errors.New("Section not found")

// ErrNameNotFound is the error returned when a config name is not found
var ErrNameNotFound = errors.New("Name not found")

type configLocation struct {
	ConfigFile string
}

type config map[string]map[string]string

// New provides a general.Config interface
func New() (general.Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filepath.Join(wd, "config-location.json"))
	if err != nil {
		return nil, err
	}
	cfgLoc := &configLocation{}
	err = json.Unmarshal(data, cfgLoc)
	if err != nil {
		return nil, err
	}

	data, err = ioutil.ReadFile(cfgLoc.ConfigFile)
	if err != nil {
		return nil, err
	}
	cfg := make(config)
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (r config) Value(section, name string) (string, error) {
	s, ok := r[section]
	if !ok {
		return "", fmt.Errorf("%w: %v", ErrSectionNotFound, section)
	}
	n, ok := s[name]
	if !ok {
		return "", fmt.Errorf("%w: %v", ErrNameNotFound, name)
	}
	return n, nil
}
