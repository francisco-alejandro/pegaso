package pegaso_test

import (
	"github.com/francisco-alejandro/pegaso"
	"github.com/stretchr/testify/assert"
	"testing"
)

const message = "This is a message"

func TestPubSub(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	bus := pegaso.CommandBus{}

	bus.Subscribe(ch1)
	bus.Subscribe(ch2)

	bus.Publish(message)

	assert.Equal(t, message, <-ch1)
	assert.Equal(t, message, <-ch2)
}
