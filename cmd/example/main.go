package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/xivdata/goxiv"
)

func main() {
	goxiv := goxiv.GoXIV{}
	char := goxiv.ScrapeCharacter(10477093)
	b, err := json.Marshal(char)
	if err != nil {
		logrus.Error("Error: ", err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("char-%d.json", char.ID), b, 0644)
	if err != nil {
		logrus.Error("Could not write json ", err)
	}
	fc := goxiv.ScrapeFreecompany(9232801448574584889)
	b, err = json.Marshal(fc)
	if err != nil {
		logrus.Error("Error: ", err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("fc-%d.json", fc.ID), b, 0644)
	if err != nil {
		logrus.Error("Could not write json ", err)
	}
	pvpteam := goxiv.ScrapePvPTeam("50276fadbb2edce09708ed5171a93c2d05eaf701")
	b, err = json.Marshal(pvpteam)
	if err != nil {
		logrus.Error("Error: ", err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("pvpteam-%v.json", pvpteam.ID), b, 0644)
	if err != nil {
		logrus.Error("Could not write json ", err)
	}
	worldlinkshell := goxiv.ScrapeLinkshell("09fc154c707570cf2a3e12f48aff36ea2506e88c", true)
	b, err = json.Marshal(worldlinkshell)
	if err != nil {
		logrus.Error("Error: ", err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("worldlinkshell-%v.json", worldlinkshell.ID), b, 0644)
	if err != nil {
		logrus.Error("Could not write json ", err)
	}
	linkshell := goxiv.ScrapeLinkshell("18858823439663593", false)
	b, err = json.Marshal(linkshell)
	if err != nil {
		logrus.Error("Error: ", err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("linkshell-%v.json", linkshell.ID), b, 0644)
	if err != nil {
		logrus.Error("Could not write json ", err)
	}
}
