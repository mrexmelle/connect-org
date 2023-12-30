package opts

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "connect-orgs",
	Short:   "Connect Orgs",
	Long:    "Connect Orgs - An organization service for Connect",
	Version: "0.1.0",
}
