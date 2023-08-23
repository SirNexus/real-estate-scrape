package realestatescrape

import (
	"log"

	"github.com/SirNexus/real-estate-scrape/internal/cmd/realestatescrape/profiles"
	"github.com/spf13/cobra"
)

func Main() {
	if err := rootCommand().Execute(); err != nil {
		log.Fatal(err.Error())
	}
}

func rootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "real-estate-scrape [command]",
	}

	cmd.AddCommand(profiles.ProfilesCommand())

	return cmd
}
