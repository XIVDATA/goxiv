package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/xivdata/goxiv/controller"
)

func main() {
	char := controller.ScrapeCharacter(10477093)
	b, err := json.Marshal(char)
	if err != nil {
		logrus.Error("Error: ", err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("char-%d.json", char.ID), b, 0644)
	if err != nil {
		logrus.Error("Could not write json ", err)
	}
	controller.ScrapeFreecompany(9232801448574584889)
}
