package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func providerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider",
		Short: "interact with providers",
	}
	cmd.AddCommand(providerListCmd())

	proCmd := prov.Cmd()
	proCmd.Use = provName
	cmd.AddCommand(proCmd)
	return cmd
}

func providerListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "print all available providers",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Available providers:", listProviders())
		},
	}
	return cmd
}

func listProviders() []string {
	keys := make([]string, len(providers))

	i := 0
	for k := range providers {
		keys[i] = k
		i++
	}
	return keys
}

func applyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apply",
		Short: "[Requires Provider] apply the configuration to the target",
	}
	cmd.Run = func(cmd *cobra.Command, args []string) {
	}
	return cmd
}

func diffCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "diff",
		Short: "[Requires Provider] print differences between the configuration and the target",
	}
	cmd.Run = func(cmd *cobra.Command, args []string) {}
	return cmd
}

func showCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "[Requires Provider] print the jsonnet in the target state format",
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		var rawDict map[string]interface{}
		raw, err := eval()
		if err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(raw), &rawDict); err != nil {
			return err
		}

		state, err := prov.Show(rawDict)
		if err != nil {
			return err
		}
		out, err := json.Marshal(state)
		if err != nil {
			return err
		}
		fmt.Println(out)
		return nil
	}
	return cmd
}
