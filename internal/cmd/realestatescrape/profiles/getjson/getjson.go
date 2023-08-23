package getjson

import (
	"fmt"

	"github.com/SirNexus/real-estate-scrape/internal/pkg/webscraperscrapeconfig"
	"github.com/spf13/cobra"
)

func AddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-json <profile-name>",
		Args:  cobra.ExactArgs(1),
		Short: "get json for configuring webscraper",
		RunE: func(cmd *cobra.Command, args []string) error {
			profileName := args[0]
			for _, profile := range webscraperscrapeconfig.Profiles {
				if profile.Name() == profileName {
					fmt.Println(profile)
					return nil
				}
			}

			return fmt.Errorf("could not find profile with name %v", profileName)
		},
	}

	return cmd
}
