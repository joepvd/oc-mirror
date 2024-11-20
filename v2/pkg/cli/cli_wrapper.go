package cli

import (
	"fmt"

	"github.com/openshift/oc-mirror/v2/internal/pkg/cli"
	clog "github.com/openshift/oc-mirror/v2/internal/pkg/log"
	"github.com/spf13/cobra"
)

/*
V2Cmd is exposed only temporary and will be removed in the future, please do not use it.
This func exists only to redirect the flow to the v2 version of the cli while v1 is still supported
*/
func V2Cmd(loglevel string) *cobra.Command {
	log := clog.New(loglevel)

	fmt.Println()

	cmd := cli.NewMirrorCmd(log)
	return cmd
}
