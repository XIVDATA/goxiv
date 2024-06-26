package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/xivdata/goxiv"
)

func main() {
	logrus.SetReportCaller(true)
	scraper := goxiv.GoXIV{}
	char := scraper.ScrapeCharacter(10477093, "jp")
	b, err := json.Marshal(char)
	if err != nil {
		logrus.Error("Error: ", err)
	}
	err = os.WriteFile(fmt.Sprintf("char-%d.json", char.ID), b, 0644)
	if err != nil {
		logrus.Error("Could not write json ", err)
	}
	fc := scraper.ScrapeFreecompany(9232801448574584889, "jp")
	b, err = json.Marshal(fc)
	if err != nil {
		logrus.Error("Error: ", err)
	}
	err = os.WriteFile(fmt.Sprintf("fc-%d.json", fc.ID), b, 0644)
	if err != nil {
		logrus.Error("Could not write json ", err)
	}
	pvpteam := scraper.ScrapePvPTeam("50276fadbb2edce09708ed5171a93c2d05eaf701", "eu")
	b, err = json.Marshal(pvpteam)
	if err != nil {
		logrus.Error("Error: ", err)
	}
	err = os.WriteFile(fmt.Sprintf("pvpteam-%v.json", pvpteam.ID), b, 0644)
	if err != nil {
		logrus.Error("Could not write json ", err)
	}
	worldlinkshell := scraper.ScrapeLinkshell("09fc154c707570cf2a3e12f48aff36ea2506e88c", true, "eu")
	b, err = json.Marshal(worldlinkshell)
	if err != nil {
		logrus.Error("Error: ", err)
	}
	err = os.WriteFile(fmt.Sprintf("worldlinkshell-%v.json", worldlinkshell.ID), b, 0644)
	if err != nil {
		logrus.Error("Could not write json ", err)
	}
	linkshell := scraper.ScrapeLinkshell("18858823439663593", false, "eu")
	b, err = json.Marshal(linkshell)
	if err != nil {
		logrus.Error("Error: ", err)
	}
	err = os.WriteFile(fmt.Sprintf("linkshell-%v.json", linkshell.ID), b, 0644)
	if err != nil {
		logrus.Error("Could not write json ", err)
	}
}
