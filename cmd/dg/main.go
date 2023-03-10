package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var rootCmd = &cobra.Command{
		Use: "dg [sub]",
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
