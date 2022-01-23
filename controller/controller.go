package controller

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	proxyfunc colly.ProxyFunc
}

func (c Controller) SetProxys(proxys ...string) {
	p, err := proxy.RoundRobinProxySwitcher(proxys...)
	if err != nil {
		logrus.Error("Error while creating roundrobin proxyswitcher: ", err)
	}
	c.proxyfunc = p
}
