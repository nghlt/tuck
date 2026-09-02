package cmd

import (
	"fmt"
	"os"

	"github.com/rot1024/tuck/session"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:     "attach [name]",
	Aliases: []string{"a"},
	Short:   "Attach to a session (creates if not exists)",
	Long: `Attach to an existing session with the given name.
If no name is specified, attaches to the most recently active session.
If the session does not exist, creates a new session and attaches to it.

Use ` + "`." + ` (default) or configured detach key to detach.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		checkNotNested()

		var name string
		if len(args) == 0 {
			// Attach to most recent session
			s, err := session.MostRecent()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			if s == nil {
				name = generateSessionName()
				createAndAttachSession(name, nil)
				return
			}
			name = s.Name
		} else {
			name = args[0]
			if !session.Exists(name) {
				createAndAttachSession(name, nil)
				return
			}
		}

		if err := session.Attach(name, session.AttachOptions{
			Quiet:       quietFlag,
			DetachKeys:  mustGetDetachKeys(),
			TitleFormat: getTitleFormat(),
		}); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}
