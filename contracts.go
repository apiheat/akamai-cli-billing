package main

import (
	"github.com/urfave/cli"
)

func cmdContractsUsage(c *cli.Context) error {
	return listContractsUsage(c)
}

func cmdContractsStatistics(c *cli.Context) error {
	return listContractsStatistics(c)
}

func listContractsUsage(c *cli.Context) error {
	return nil
}

func listContractsStatistics(c *cli.Context) error {
	return nil
}
