/*
Copyright Scoir Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/scoir/canis/pkg/controller"
	"github.com/scoir/canis/pkg/didcomm/mediator"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the didcomm mediator service",
	Long:  `Starts a didcomm mediator service`,
	Run:   runStart,
}

func runStart(_ *cobra.Command, _ []string) {
	i, err := mediator.New(ctx)
	if err != nil {
		log.Fatalln("unable to initialize mediator", err)
	}

	runner, err := controller.New(ctx, i)
	if err != nil {
		log.Fatalln("unable to start didcomm-mediator", err)
	}

	err = runner.Launch()
	if err != nil {
		log.Fatalln("launch errored with", err)
	}

}

func init() {
	rootCmd.AddCommand(startCmd)
}
