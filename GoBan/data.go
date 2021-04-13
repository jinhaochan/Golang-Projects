package main

import (
	"io/ioutil"
)

type Page map[string]string

func (p Page) saveData(activity string) error {
	data := p[activity]
	return ioutil.WriteFile(activity+".txt", []byte(data), 0600)
}

func loadAllData() *Page {

	activities := [3]string{"todo", "progress", "done"}

	PageData := make(Page)

	for _, activity := range activities {
		PageData[activity] = string(loadData(activity))
	}

	return &PageData
}

func loadData(activity string) []byte {

	data, err := ioutil.ReadFile("data/" + activity + ".txt")
	if err != nil {
		return nil
	}

	return data
}
