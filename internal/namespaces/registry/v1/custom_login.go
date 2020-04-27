package registry

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"reflect"

	"github.com/scaleway/scaleway-cli/internal/core"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type registryLoginArgs struct {
	Region  scw.Region
	Program string
}

func registryLoginCommand() *core.Command {
	return &core.Command{
		Short: `Login to a registry`,
		Long: `This command will run the correct command in order to log you in on the registry with the chosen program.
You will need to have the chosen binary installed on your system and in your PATH.`,
		Namespace: "registry",
		Resource:  "login",
		ArgsType:  reflect.TypeOf(registryLoginArgs{}),
		ArgSpecs: []*core.ArgSpec{
			{
				Name:       "program",
				Short:      "Program used to log in to the namespace",
				Default:    core.DefaultValueSetter(string(docker)),
				EnumValues: availablePrograms.StringArray(),
			},
			core.RegionArgSpec(scw.AllRegions...),
		},
		Run: registryLoginRun,
	}
}

func registryLoginRun(ctx context.Context, argsI interface{}) (i interface{}, e error) {
	args := argsI.(*registryLoginArgs)
	client := core.ExtractClient(ctx)

	region := args.Region.String()
	if region == "" {
		scwRegion, ok := client.GetDefaultRegion()
		if !ok {
			return nil, fmt.Errorf("no default region configured")
		}
		region = scwRegion.String()
	}
	endpoint := endpointPrefix + region + endpointSuffix

	secretKey, ok := client.GetSecretKey()
	if !ok {
		return nil, fmt.Errorf("could not get secret key")
	}

	var cmdArgs []string

	switch program(args.Program) {
	case docker, podman:
		cmdArgs = []string{"login", "-u", "scaleway", "--password-stdin", endpoint}
	default:
		return nil, fmt.Errorf("unknown program")
	}

	cmd := exec.Command(args.Program, cmdArgs...)
	cmd.Stdin = bytes.NewBufferString(secretKey)

	exitCode, err := core.ExecCmd(ctx, cmd)

	if err != nil {
		return nil, err
	}
	if exitCode != 0 {
		return nil, &core.CliError{Empty: true, Code: exitCode}
	}

	return &core.SuccessResult{
		Empty: true, // the program will output the success message
	}, nil
}
