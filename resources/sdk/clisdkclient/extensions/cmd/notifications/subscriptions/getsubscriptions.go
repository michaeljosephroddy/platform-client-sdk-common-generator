package subscriptions

import (
	"fmt"
	"gc/utils"
	"gc/services"
	"log"

	"github.com/spf13/cobra"
)

var getSubscriptionsCmd = &cobra.Command{
	Use:   "get [channel id]",
	Short: "Retrieves a list of subscriptions for a channel",
	Long:  `Retrieves a list of subscriptions for a channel`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		targetURI := fmt.Sprintf("%s/channels/%s/subscriptions", BaseURI, args[0])
		retryFunc := services.Retry(targetURI, CommandService.Get)
		results, err := retryFunc(nil)
		if err != nil {
			log.Fatal(err)
		}

		utils.Render(results)
	},
}
