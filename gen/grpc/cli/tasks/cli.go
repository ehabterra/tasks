// Code generated by goa v3.2.3, DO NOT EDIT.
//
// tasks gRPC client CLI support package
//
// Command:
// $ goa gen tasks/design

package cli

import (
	"flag"
	"fmt"
	"os"
	tasksc "tasks/gen/grpc/tasks/client"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
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
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		tasksFlags = flag.NewFlagSet("tasks", flag.ContinueOnError)

		tasksListFlags    = flag.NewFlagSet("list", flag.ExitOnError)
		tasksListViewFlag = tasksListFlags.String("view", "", "")

		tasksShowFlags       = flag.NewFlagSet("show", flag.ExitOnError)
		tasksShowMessageFlag = tasksShowFlags.String("message", "", "")
		tasksShowViewFlag    = tasksShowFlags.String("view", "", "")

		tasksAddFlags       = flag.NewFlagSet("add", flag.ExitOnError)
		tasksAddMessageFlag = tasksAddFlags.String("message", "", "")

		tasksUpdateFlags       = flag.NewFlagSet("update", flag.ExitOnError)
		tasksUpdateMessageFlag = tasksUpdateFlags.String("message", "", "")

		tasksRemoveFlags       = flag.NewFlagSet("remove", flag.ExitOnError)
		tasksRemoveMessageFlag = tasksRemoveFlags.String("message", "", "")

		tasksStatusFlags       = flag.NewFlagSet("status", flag.ExitOnError)
		tasksStatusMessageFlag = tasksStatusFlags.String("message", "", "")
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
			c := tasksc.NewClient(cc, opts...)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = tasksc.BuildListPayload(*tasksListViewFlag)
			case "show":
				endpoint = c.Show()
				data, err = tasksc.BuildShowPayload(*tasksShowMessageFlag, *tasksShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = tasksc.BuildAddPayload(*tasksAddMessageFlag)
			case "update":
				endpoint = c.Update()
				data, err = tasksc.BuildUpdatePayload(*tasksUpdateMessageFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = tasksc.BuildRemovePayload(*tasksRemoveMessageFlag)
			case "status":
				endpoint = c.Status()
				data, err = tasksc.BuildStatusPayload(*tasksStatusMessageFlag)
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
	fmt.Fprintf(os.Stderr, `%s [flags] tasks show -message JSON -view STRING

Show task by ID
    -message JSON: 
    -view STRING: 

Example:
    `+os.Args[0]+` tasks show --message '{
      "id": "Eos consequatur exercitationem necessitatibus quae."
   }' --view "tiny"
`, os.Args[0])
}

func tasksAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks add -message JSON

Add new task and return ID.
    -message JSON: 

Example:
    `+os.Args[0]+` tasks add --message '{
      "assignee": {
         "email": "ehabterra@hotmail.com",
         "firstname": "Ehab",
         "isactive": false,
         "lastname": "Terra",
         "role": "admin"
      },
      "created_date": "2008-08-22T21:23:49Z",
      "description": "Task description",
      "due_date": "2011-10-03T01:36:01Z",
      "owner": {
         "email": "ehabterra@hotmail.com",
         "firstname": "Ehab",
         "isactive": false,
         "lastname": "Terra",
         "role": "admin"
      },
      "status": "Closed",
      "title": "New task title",
      "updated_date": "1975-06-12T01:40:28Z"
   }'
`, os.Args[0])
}

func tasksUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks update -message JSON

Update existing task and return ID.
    -message JSON: 

Example:
    `+os.Args[0]+` tasks update --message '{
      "id": "Consectetur explicabo fugit tenetur.",
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
   }'
`, os.Args[0])
}

func tasksRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks remove -message JSON

Remove task from tasks data
    -message JSON: 

Example:
    `+os.Args[0]+` tasks remove --message '{
      "id": "Quis deserunt deleniti dolore."
   }'
`, os.Args[0])
}

func tasksStatusUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks status -message JSON

change task status by id
    -message JSON: 

Example:
    `+os.Args[0]+` tasks status --message '{
      "id": "Voluptas ad.",
      "status": "Closed"
   }'
`, os.Args[0])
}
