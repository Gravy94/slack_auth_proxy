package main

import (
	"net/url"
)

type UpstreamConfiguration struct {
	Host 	string	 `yaml:"host,omitempty"`
	HostURL *url.URL `yaml:"-"`

	Users 	[]string `yaml:"users,omitempty"`
	Groups  []string `yaml:"groups,omitempty"`
}

type UpstreamConfigurationMap map[string]*UpstreamConfiguration

func (c *UpstreamConfiguration) Parse() (err error) {
	c.HostURL, err = url.Parse(c.Host)

	return
}