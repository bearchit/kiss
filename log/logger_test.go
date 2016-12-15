package log_test

import (
	"testing"

	"github.com/bearchit/kiss/log"
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	logger := log.New()

	assert.NotNil(t, logger)
}
