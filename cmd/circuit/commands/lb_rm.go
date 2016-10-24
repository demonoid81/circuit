package commands

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var lbDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a service",
	Long: `Delete a service
Example:
    circuit lb delete <name>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			return
		}

		name := args[0]

		if name == "" {
			cmd.Help()
			logrus.Fatal("you must specify a service name")
		}

		c, err := getController(cmd)
		if err != nil {
			logrus.Fatal(err)
		}

		if err := c.RemoveService(name); err != nil {
			logrus.Fatalf("error removing service: %s", err)
		}

		logrus.Infof("service %s removed", name)
	},
}