package main

import (
	"github.com/urfave/cli"
)

func cmdRGsUsage(c *cli.Context) error {
	return listRGsUsage(c)
}

func cmdRGsStatistics(c *cli.Context) error {
	return listRGsStatistics(c)
}

func listRGsUsage(c *cli.Context) error {
	return nil
}

func listRGsStatistics(c *cli.Context) error {
	return nil
}
