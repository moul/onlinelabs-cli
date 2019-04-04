/*
 * Minio Client (C) 2018 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"github.com/fatih/color"
	"github.com/minio/cli"
	"github.com/minio/mc/pkg/console"
	"github.com/minio/mc/pkg/probe"
	"github.com/minio/minio/pkg/madmin"
)

var adminUsersDisableCmd = cli.Command{
	Name:   "disable",
	Usage:  "Disable users",
	Action: mainAdminUsersDisable,
	Before: setGlobalsFromContext,
	Flags:  globalFlags,
	CustomHelpTemplate: `NAME:
  {{.HelpName}} - {{.Usage}}

USAGE:
  {{.HelpName}} TARGET USERNAME

FLAGS:
  {{range .VisibleFlags}}{{.}}
  {{end}}
EXAMPLES:
  1. Disable a user 'newuser' for Minio server.
     $ {{.HelpName}} myminio newuser
`,
}

// checkAdminUsersDisableSyntax - validate all the passed arguments
func checkAdminUsersDisableSyntax(ctx *cli.Context) {
	if len(ctx.Args()) != 2 {
		cli.ShowCommandHelpAndExit(ctx, "disable", 1) // last argument is exit code
	}
}

// mainAdminUsersDisable is the handle for "mc admin users disable" command.
func mainAdminUsersDisable(ctx *cli.Context) error {
	checkAdminUsersDisableSyntax(ctx)

	console.SetColor("UserMessage", color.New(color.FgGreen))

	// Get the alias parameter from cli
	args := ctx.Args()
	aliasedURL := args.Get(0)

	// Create a new Minio Admin Client
	client, err := newAdminClient(aliasedURL)
	fatalIf(err, "Cannot get a configured admin connection.")

	e := client.SetUserStatus(args.Get(1), madmin.AccountDisabled)
	fatalIf(probe.NewError(e).Trace(args...), "Cannot disable user")

	printMsg(userMessage{
		op:        "disable",
		AccessKey: args.Get(1),
	})

	return nil
}
