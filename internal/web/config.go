package web

import "fmt"

type WebConfig struct {
	Host string
	Port string
}

func (c WebConfig) Address() string {
	return fmt.Sprintf("%s%s", c.Host, c.Port)
}
