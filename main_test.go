package main

import (
	"testing"

	"github.com/sensu/sensu-go/types"
	"github.com/stretchr/testify/assert"
)

func TestSendMetrics(t *testing.T) {
	assert := assert.New(t)
	event := types.FixtureEvent("entity1", "check1")
	event.Check = nil
	event.Metrics = types.FixtureMetrics()

	err := sendMetrics(event)
	assert.NoError(err)
}
