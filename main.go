package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main(){
	var rootCmd = &cobra.Command{
		Use: "midnight",
		Short: "high-performance security audit cli for solo builders",
	}	

	var prelaunchCmd = &cobra.Command{
		Use: "prelaunch",
		Short: "scan your entire codebase before merging or deploying",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string){
			fmt.Println("prelaunch scan:", ".")
		},

	}

	var productionCmd = &cobra.Command{
		Use: "production",
		Short:"deep audit for production readiness",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string){
			fmt.Println("production scan:",".")
		},
	}

	var quickscanCmd = &cobra.Command{
		Use:   "quickscan [file]",
		Short: "scan a single file instantly",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("quickscan:", args[0])
		},
	}

	rootCmd.AddCommand(prelaunchCmd, quickscanCmd, productionCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}