package main

import (
	"encoding/json"
	"io/ioutil"
)

// ConfigStore type for storing Entries
type ConfigStore interface {
	storeEntries(*[]Entry) error
	loadEntries() (*[]Entry, error)
}

//JSONConfigStore stores entries in json file
type JSONConfigStore struct{}

func (JSONConfigStore) storeEntries(entries *[]Entry) error {
	encoded, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(configFileName, encoded, 0644); err != nil {
		return err
	}
	return nil
}

func (JSONConfigStore) loadEntries() (entries *[]Entry, err error) {
	encoded, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return
	}
	if err = json.Unmarshal(encoded, &entries); err != nil {
		return &[]Entry{}, err
	}
	return
}
