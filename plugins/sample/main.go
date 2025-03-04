// +build plugin

/*
Copyright © 2019 Portworx

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"github.com/portworx/px/pkg/util"
	"github.com/spf13/cobra"
)

var (
	argExample string
)

// sampleCmd represents the get command
var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		util.Printf("Got called with example = [%s]\n", argExample)
		return nil
	},
}

// Plugin Version
var Version = "0.0.1"

// PluginInfo sets information about the plugin
var PluginManifest = map[string]string{
	"name":        "sample",
	"description": "Sample plugin",
	"version":     Version,
}

func PluginInit(parent *cobra.Command) {
	parent.AddCommand(sampleCmd)
	sampleCmd.Flags().StringVar(&argExample, "example", "", "Help message for sample")
}
