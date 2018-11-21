package main

import (
	"fmt"
	"os"
	"sort"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

var (
	apiClient       *edgegrid.Client
	appName, appVer string
)

func main() {
	app := common.CreateNewApp(appName, "A CLI to interact with Akamai Adaptive Acceleration", appVer)
	app.Flags = common.CreateFlags()

	app.Commands = []cli.Command{
		{
			Name:      "measures",
			Aliases:   []string{"c"},
			UsageText: fmt.Sprintf("%s measures", appName),
			Usage:     "List usage data matching any criteria specified by a Query object. The query needs to specify at least one contractID or reporting-groupID value",
			Action:    cmdMeasures,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "month",
					Value: 0,
					Usage: "The month (1–12) for which to aggregate data",
				},
				cli.IntFlag{
					Name:  "year",
					Value: 0,
					Usage: "The year for which to aggregate data",
				},
				cli.StringSliceFlag{
					Name:  "contractIDs",
					Usage: "Filter data to the specified set of contract identifiers, at least one item if specified. Otherwise if omitted, you need to specify a set of reporting-groupIDs",
				},
				cli.StringSliceFlag{
					Name:  "reporting-groupIDs",
					Usage: "Filter data to the specified set of reporting group identifiers, at least one item if specified. Otherwise if omitted, you need to specify a set of contractIDs",
				},
				cli.StringSliceFlag{
					Name:  "productIDs",
					Usage: "Filter data to the specified set of product identifiers. An empty array produces unfiltered results, the same as omitting the member",
				},
				cli.StringSliceFlag{
					Name:  "statistic-types",
					Usage: "Statistics for which you want to retrieve data, each corresponding to a Statistic’s name value. An empty array produces unfiltered results, the same as omitting the member",
				},
			},
		},
		{
			Name:  "reporting-groups",
			Usage: "Reporting Groups related commands. List Usage per reporting group and List Statistics per reporting group",
			Subcommands: []cli.Command{
				{
					Name:      "usage",
					UsageText: fmt.Sprintf("%s reporting-groups usage REPORTING_GROUP_ID", appName),
					Usage:     "This operation lists a period of usage for a reporting group. Reporting groups collect sets of CP codes under a contract, allowing you to track billing for related types of traffic rather than for the entire contract. Usage is aggregated daily, with one value per statistic. Unless final is true, data may still update to reflect latent edge server traffic",
					Action:    cmdRGsUsage,
				},
				{
					Name:      "statistics",
					Usage:     "Optionally run this operation if there’s only one type of information you want to include in a reporting group’s usage report, rather than the complete set. Availability of reportable statistics may depend on the term of the contract, hence the need to specify them for a range of time",
					UsageText: fmt.Sprintf("%s reporting-groups statistics REPORTING_GROUP_ID", appName),
					Action:    cmdRGsStatistics,
				},
			},
		},
		{
			Name:  "contracts",
			Usage: "Contract related commands. List Usage per contract and List Statistics per contract",
			Subcommands: []cli.Command{
				{
					Name:      "usage",
					UsageText: fmt.Sprintf("%s contracts usage CONTRACT_ID", appName),
					Usage:     "This operation lists a period of usage for an entire contract",
					Action:    cmdContractsUsage,
				},
				{
					Name:      "statistics",
					Usage:     "Optionally run this operation if there’s only one type of information you want to include in your usage report, rather than the complete set. This lists each type of reportable statistic. Availability may depend on the term of the contract, hence the need to specify them for a range of time",
					UsageText: fmt.Sprintf("%s contracts statistics CONTRACT_ID", appName),
					Action:    cmdContractsStatistics,
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Before = func(c *cli.Context) error {
		var err error

		apiClient, err = common.EdgeClientInit(c.GlobalString("config"), c.GlobalString("section"), c.GlobalString("debug"))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
