package profiles

import (
	"github.com/SirNexus/real-estate-scrape/internal/cmd/realestatescrape/profiles/getjson"
	"github.com/SirNexus/real-estate-scrape/internal/cmd/realestatescrape/profiles/list"
	"github.com/spf13/cobra"
)

func ProfilesCommand() *cobra.Command {
	profileCmd := &cobra.Command{
		Use:   "profiles [command]",
		Short: "Run operations on scrape profiles",
	}

	profileCmd.AddCommand(list.AddCommand())
	profileCmd.AddCommand(getjson.AddCommand())

	return profileCmd
}
