package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func GetAssertor(t *testing.T) *require.Assertions {

	return require.New(t)
}