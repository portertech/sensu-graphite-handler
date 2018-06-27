package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"testing"

	"github.com/sensu/sensu-go/types"
	"github.com/stretchr/testify/assert"
)

func TestSendMetrics(t *testing.T) {
	assert := assert.New(t)
	event := types.FixtureEvent("entity1", "check1")
	event.Check = nil
	event.Metrics = types.FixtureMetrics()

	host = "127.0.0.1"
	port = 2003

	go func() {
		err := sendMetrics(event)
		assert.NoError(err)
	}()

	listener, err := net.Listen("tcp", ":2003")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		buffer, err := ioutil.ReadAll(conn)
		if err != nil {
			t.Fatal(err)
		}

		point := event.Metrics.Points[0]
		stringTimestamp := strconv.FormatInt(point.Timestamp, 10)

		expected := fmt.Sprintf("%s %v %s\n", point.Name, point.Value, stringTimestamp[:10])

		if msg := string(buffer[:]); msg != expected {
			t.Fatalf("Unexpected message:\nGot:\t\t%s\nExpected:\t%s\n", msg, expected)
		}

		return
	}

}
