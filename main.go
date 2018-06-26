package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/sensu/sensu-go/types"
	"github.com/spf13/cobra"
)

var (
	host  string
	port  int
	stdin *os.File
)

func main() {
	rootCmd := configureRootCommand()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}

func configureRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sensu-graphite-handler",
		Short: "a graphite handler built for use with sensu",
		RunE:  run,
	}

	cmd.Flags().StringVarP(&host,
		"host",
		"h",
		"127.0.0.1",
		"the graphite carbon host")

	cmd.Flags().IntVarP(&port,
		"port",
		"p",
		2003,
		"the graphite carbon tcp port")

	_ = cmd.MarkFlagRequired("host")
	_ = cmd.MarkFlagRequired("port")

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		_ = cmd.Help()
		return errors.New("invalid argument(s) received")
	}

	if stdin == nil {
		stdin = os.Stdin
	}

	eventJSON, err := ioutil.ReadAll(stdin)
	if err != nil {
		return fmt.Errorf("failed to read stdin: %s", err.Error())
	}

	event := &types.Event{}
	err = json.Unmarshal(eventJSON, event)
	if err != nil {
		return fmt.Errorf("failed to unmarshal stdin data: %s", err.Error())
	}

	if err = event.Validate(); err != nil {
		return fmt.Errorf("failed to validate event: %s", err.Error())
	}

	if !event.HasMetrics() {
		return fmt.Errorf("event does not contain metrics")
	}

	return sendMetrics(event)
}

func sendMetrics(event *types.Event) error {
	var buffer bytes.Buffer

	for _, point := range event.Metrics.Points {
		stringTimestamp := strconv.FormatInt(point.Timestamp, 10)

		line := fmt.Sprintf("%s %v %s\n", point.Name, point.Value, stringTimestamp[:10])

		buffer.WriteString(line)
	}

	fmt.Println(buffer.String())

	return nil
}
