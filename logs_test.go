package lex_test

import (
	"errors"
	"testing"

	"github.com/daqiancode/lex"
)

func TestLog(t *testing.T) {
	log := lex.GetLogger("", "error", true)
	log.Info().Msg("hello")
	log.Error().Msg("error occur")
	log.Error().Stack().Err(errors.New("error info")).Msg("")
}
