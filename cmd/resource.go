package cmd

import (
	"github.com/ZhangSetSail/OChc/pkg/action"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"io"
)

var resourceLong = `helm resource ------`

func newResourceCmd(out io.Writer) *cobra.Command {
	resource := action.NewResource()
	cmd := &cobra.Command{
		Use:   "resource ",
		Short: "generate a helm chart package with one click through k8s cluster resources",
		Long:  resourceLong,
		Args:  action.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			resource.Run(args)
			return
		},
	}
	addResourceFlags(cmd, cmd.Flags(), resource)
	return cmd
}

func addResourceFlags(cmd *cobra.Command, f *pflag.FlagSet, resource *action.Resource) {
	f.StringVar(&resource.Namespace, "namespace", "", "specify template used to name the release")
	err := cmd.MarkFlagRequired("namespace")
	if err != nil {
		logrus.Errorf("ochc resource MarkFlagRequired failure: %v", err)
	}
}
