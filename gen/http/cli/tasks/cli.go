// Code generated by goa v3.2.3, DO NOT EDIT.
//
// tasks HTTP client CLI support package
//
// Command:
// $ goa gen tasks/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	tasksc "tasks/gen/http/tasks/client"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `tasks (list|show|add|update|remove|status)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` tasks list --view "default"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		tasksFlags = flag.NewFlagSet("tasks", flag.ContinueOnError)

		tasksListFlags    = flag.NewFlagSet("list", flag.ExitOnError)
		tasksListViewFlag = tasksListFlags.String("view", "", "")

		tasksShowFlags    = flag.NewFlagSet("show", flag.ExitOnError)
		tasksShowIDFlag   = tasksShowFlags.String("id", "REQUIRED", "ID of task to show")
		tasksShowViewFlag = tasksShowFlags.String("view", "", "")

		tasksAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		tasksAddBodyFlag = tasksAddFlags.String("body", "REQUIRED", "")

		tasksUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		tasksUpdateBodyFlag = tasksUpdateFlags.String("body", "REQUIRED", "")
		tasksUpdateIDFlag   = tasksUpdateFlags.String("id", "REQUIRED", "ID of task to show")

		tasksRemoveFlags  = flag.NewFlagSet("remove", flag.ExitOnError)
		tasksRemoveIDFlag = tasksRemoveFlags.String("id", "REQUIRED", "ID of task to remove")

		tasksStatusFlags    = flag.NewFlagSet("status", flag.ExitOnError)
		tasksStatusBodyFlag = tasksStatusFlags.String("body", "REQUIRED", "")
	)
	tasksFlags.Usage = tasksUsage
	tasksListFlags.Usage = tasksListUsage
	tasksShowFlags.Usage = tasksShowUsage
	tasksAddFlags.Usage = tasksAddUsage
	tasksUpdateFlags.Usage = tasksUpdateUsage
	tasksRemoveFlags.Usage = tasksRemoveUsage
	tasksStatusFlags.Usage = tasksStatusUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "tasks":
			svcf = tasksFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "tasks":
			switch epn {
			case "list":
				epf = tasksListFlags

			case "show":
				epf = tasksShowFlags

			case "add":
				epf = tasksAddFlags

			case "update":
				epf = tasksUpdateFlags

			case "remove":
				epf = tasksRemoveFlags

			case "status":
				epf = tasksStatusFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "tasks":
			c := tasksc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = tasksc.BuildListPayload(*tasksListViewFlag)
			case "show":
				endpoint = c.Show()
				data, err = tasksc.BuildShowPayload(*tasksShowIDFlag, *tasksShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = tasksc.BuildAddPayload(*tasksAddBodyFlag)
			case "update":
				endpoint = c.Update()
				data, err = tasksc.BuildUpdatePayload(*tasksUpdateBodyFlag, *tasksUpdateIDFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = tasksc.BuildRemovePayload(*tasksRemoveIDFlag)
			case "status":
				endpoint = c.Status()
				data, err = tasksc.BuildStatusPayload(*tasksStatusBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// tasksUsage displays the usage of the tasks command and its subcommands.
func tasksUsage() {
	fmt.Fprintf(os.Stderr, `The tasks service performs task data.
Usage:
    %s [globalflags] tasks COMMAND [flags]

COMMAND:
    list: List all stored tasks
    show: Show task by ID
    add: Add new task and return ID.
    update: Update existing task and return ID.
    remove: Remove task from tasks data
    status: change task status by id

Additional help:
    %s tasks COMMAND --help
`, os.Args[0], os.Args[0])
}
func tasksListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks list -view STRING

List all stored tasks
    -view STRING: 

Example:
    `+os.Args[0]+` tasks list --view "default"
`, os.Args[0])
}

func tasksShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks show -id STRING -view STRING

Show task by ID
    -id STRING: ID of task to show
    -view STRING: 

Example:
    `+os.Args[0]+` tasks show --id "Velit nesciunt totam autem quos quia." --view "default"
`, os.Args[0])
}

func tasksAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks add -body JSON

Add new task and return ID.
    -body JSON: 

Example:
    `+os.Args[0]+` tasks add --body '{
      "assignee": {
         "email": "ehabterra@hotmail.com",
         "firstname": "Ehab",
         "isactive": false,
         "lastname": "Terra",
         "role": "admin"
      },
      "created_date": "1972-12-11T14:31:10Z",
      "description": "Task description",
      "due_date": "2001-09-13T00:38:25Z",
      "owner": {
         "email": "ehabterra@hotmail.com",
         "firstname": "Ehab",
         "isactive": false,
         "lastname": "Terra",
         "role": "admin"
      },
      "status": "Closed",
      "title": "New task title",
      "updated_date": "2008-01-17T02:07:49Z"
   }'
`, os.Args[0])
}

func tasksUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks update -body JSON -id STRING

Update existing task and return ID.
    -body JSON: 
    -id STRING: ID of task to show

Example:
    `+os.Args[0]+` tasks update --body '{
      "task": {
         "assignee": {
            "email": "ehabterra@hotmail.com",
            "firstname": "Ehab",
            "isactive": false,
            "lastname": "Terra",
            "role": "admin"
         },
         "created_date": "1977-08-23T17:31:33Z",
         "description": "Task description",
         "due_date": "1994-02-04T20:22:05Z",
         "id": "Dolor eos dolorem numquam odio aspernatur et.",
         "owner": {
            "email": "ehabterra@hotmail.com",
            "firstname": "Ehab",
            "isactive": false,
            "lastname": "Terra",
            "role": "admin"
         },
         "status": "Pending",
         "title": "New task title",
         "updated_date": "1994-11-14T18:35:02Z"
      }
   }' --id "Dolorem nulla nihil nihil."
`, os.Args[0])
}

func tasksRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks remove -id STRING

Remove task from tasks data
    -id STRING: ID of task to remove

Example:
    `+os.Args[0]+` tasks remove --id "Fugit amet."
`, os.Args[0])
}

func tasksStatusUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks status -body JSON

change task status by id
    -body JSON: 

Example:
    `+os.Args[0]+` tasks status --body '{
      "id": "Qui dignissimos.",
      "status": "Open"
   }'
`, os.Args[0])
}