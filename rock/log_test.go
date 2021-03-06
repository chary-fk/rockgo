package rock

import (
	"testing"

	"github.com/byte-power/rockgo/log"
	"github.com/byte-power/rockgo/util"
	"github.com/stretchr/testify/assert"
)

func TestFluentLog(t *testing.T) {
	config := make(util.AnyMap)
	config["level"] = "info"
	config["tag"] = "test-fluent"
	config["port"] = 24225

	_, err := parseFluentLogger(config)
	if err == nil {
		t.Error("Should error if host not in config.")
	}

	config["host"] = "127.0.0.1"
	config["async"] = true
	logger, _ := parseFluentLogger(config)

	if logger == nil {
		t.Error("Create FluentLogger failed.")
	}
}

func TestParseLogComponents(t *testing.T) {
	assert.Equal(t, log.LevelDebug, parseLevel(nil))
	assert.Equal(t, log.LevelDebug, parseLevel("Debug"))
	assert.Equal(t, log.LevelInfo, parseLevel("inFo"))
	assert.Equal(t, log.LevelWarn, parseLevel("Warn"))
	assert.Equal(t, log.LevelError, parseLevel("erroR"))
	assert.Equal(t, log.LevelFatal, parseLevel("fatAl"))
	assert.Equal(t, log.MessageFormatJSON, parseMessageFormat("json"))
	assert.Equal(t, log.MessageFormatText, parseMessageFormat("text"))
	assert.Equal(t, log.TimeFormatISO8601, log.MakeTimeFormat("iso8601"))
	assert.Equal(t, log.TimeFormatSeconds, log.MakeTimeFormat("seconds"))
	assert.Equal(t, log.TimeFormatMillis, log.MakeTimeFormat("millis"))
	assert.Equal(t, log.TimeFormatNanos, log.MakeTimeFormat("nanos"))
}
