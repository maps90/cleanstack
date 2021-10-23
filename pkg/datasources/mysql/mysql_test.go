// Package mysql storage is designed to give lazy load singleton access to mysql connections
// it doesn't provide any cluster nor balancing support, assuming it is handled
// in lower level infra, i.e. proxy, cluster etc.
package mysql

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	require := require.New(t)
	err := Init()
	require.Nil(err)
}
