package cmd

import (
	"fmt"
	"io/ioutil"
)

func (c *Cmd) Write() error {
	if len(c.Name) == 0 {
		return fmt.Errorf("Name is empty")
	} else if c.Name[0] != '/' {
		return fmt.Errorf("Name must start with a '/'")
	}

	if args := c.Args(); len(args) != 0 {
		return fmt.Errorf("Write does not accept arguments")
	}

	fp, _ := c.Path()
	if err := ioutil.WriteFile(fp, []byte(c.Text), 0644); err != nil {
		return err
	}
	c.Clean = true
	return nil
}
