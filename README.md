# UNDER HEAVY DEVELOPMENT

## Akamai CLI for Billing Center

The Billing Center API allows you to access your website’s monthly usage over the Akamai network. Usage data is available for entire contracts, or for reporting groups, which may include CP codes from more than one contract.

Should you miss something we *gladly accept patches* :)

CLI uses custom [Akamai API client](https://github.com/apiheat/go-edgegrid)

## Configuration & Installation

### Credentials

Set up your credential files as described in the [authorization](https://developer.akamai.com/introduction/Prov_Creds.html) and [credentials](https://developer.akamai.com/introduction/Conf_Client.html) sections of the getting started guide on developer.akamai.com.

Tools expect proper format of sections in edgerc file which example is shown below

*NOTE:* Default file location is *~/.edgerc*

```
[default]
client_secret = XXXXXXXXXXXX
host = XXXXXXXXXXXX
access_token = XXXXXXXXXXXX
client_token = XXXXXXXXXXXX
```

In order to change section which is being actively used you can

* change it via `--config parameter` of the tool itself
* change it via env variable `export AKAMAI_EDGERC_CONFIG=/Users/jsmitsh/.edgerc`

In order to change section which is being actively used you can

* change it via `--section parameter` of the tool itself
* change it via env variable `export AKAMAI_EDGERC_SECTION=mycustomsection`

> *NOTE:* Make sure your API client do have appropriate scopes enabled

### Installation

The tool can be used as a stand-alone binary or in conjuction with [Akamai CLI](https://developer.akamai.com/cli).

#### Akamai-cli ( recommended )

Execute the following from console

```shell
> akamai install https://github.com/apiheat/akamai-cli-billing
```

#### Stand-alone

As part of automated releases/builds you can download latest version from the project release page

## Usage

```shell
NAME:
   akamai-cli-billing - A CLI to interact with Akamai Adaptive Acceleration

USAGE:
   akamai-cli-billing [global options] command [command options] [arguments...]

VERSION:
   X.X.X

AUTHORS:
   Petr Artamonov
   Rafal Pieniazek

COMMANDS:
     contracts         Contract related commands. List Usage per contract and List Statistics per contract
     measures, c       List usage data matching any criteria specified by a Query object. The query needs to specify at least one contractID or reporting-groupID value
     reporting-groups  Reporting Groups related commands. List Usage per reporting group and List Statistics per reporting group
     help, h           Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE   Location of the credentials FILE (default: "/Users/USERNAME/.edgerc") [$AKAMAI_EDGERC_CONFIG]
   --debug value            Debug Level [$AKAMAI_EDGERC_DEBUGLEVEL]
   --section NAME, -s NAME  NAME of section to use from credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
   --help, -h               show help
   --version, -v            print the version
```

## Development

In order to develop the tool with us do the following:

1. Fork repository
1. Clone it to your folder ( within *GO* path )
1. Ensure you can restore dependencies by running

   ```shell
   dep ensure
   ```

1. Make necessary changes
1. Make sure solution builds properly ( feel free to add tests )

   ```shell
   go build -ldflags="-s -w -X main.appVer=1.2.3 -X main.appName=$(basename `pwd`)"
   ```
