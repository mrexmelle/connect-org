package opts

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "connect-org",
	Short:   "Connect Org",
	Long:    "Connect Org - An organization service for Connect",
	Version: "0.2.0",
}
