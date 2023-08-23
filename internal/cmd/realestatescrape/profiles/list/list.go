package list

import (
	"fmt"

	"github.com/SirNexus/real-estate-scrape/internal/pkg/webscraperscrapeconfig"
	"github.com/spf13/cobra"
)

func AddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all profiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, profile := range webscraperscrapeconfig.Profiles {
				fmt.Println(profile.Name())
			}

			return nil
		},
	}

	return cmd
}
