package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vrypan/fargo/localdb"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
	Run:   cacheManager,
}

func cacheManager(cmd *cobra.Command, args []string) {
	localdb.Open()
	defer localdb.Close()
	max, top, bottom := localdb.Stats()
	fmt.Println("No actual cache management yet...")
	fmt.Println("If ~/.fargo/local.db gets too big, just remove it, and it will be re-created from zero.")
	fmt.Printf("Top: %v\n", top)
	fmt.Printf("Max: %v\n", max)
	fmt.Printf("Bottom: %v\n", bottom)
}

func init() {
	rootCmd.AddCommand(cacheCmd)
}
