package lex_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/daqiancode/lex"
)

func TestRollingLog(t *testing.T) {
	fmt.Println(lex.GetLocalIPs())
	rw := lex.NewRollingWriter("lex.log", 10, 3, 3, true, false)
	tags := map[string]string{"app": "lex"}
	log := lex.NewLogger(tags, rw, os.Stdout)
	log.Info().Msg("hello")
	log.Error().Msg("error occur")
	log.Error().Stack().Err(errors.New("error info")).Send()

}
