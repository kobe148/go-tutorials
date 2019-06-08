package test

import (
	"strings"

	"testing"
	
	"github.com/stretchr/testify/require"

	"hw/biglog" // You need to implement a biglog package.
)

// In this challenge, you need to implement a log system. It basically has 2 API, one for log and another for search.
func TestChallenge1(t *testing.T) {
	server := biglog.StartServer() // starts a server in its own goroutine and return the server.
	defer func() {
		err := server.Close()
		require.NoError(t, err)
	}()

	// The only requirement is that it accepts an io.Reader. You need to decide the storage mechanism.
	err := server.Log(strings.NewReader(`{"time": "2019-01-06", "number": 123}`)); require.NoError(t, err)
	err = server.Log(strings.NewReader(`{"time": "2019-01-05", "number": 124}`)); require.NoError(t, err)
	err = server.Log(strings.NewReader(`{"time": "2019-01-04", "number": 125}`)); require.NoError(t, err)
	err = server.Log(strings.NewReader(`{"time": "2019-01-03", "number": 126}`)); require.NoError(t, err)
	err = server.Log(strings.NewReader(`{"time": "2019-01-02", "number": 127}`)); require.NoError(t, err)
	err = server.Log(strings.NewReader(`{"time": "2019-01-01", "number": 128}`)); require.NoError(t, err)

	// You should come out a search API.
	// Find the record between 01-01 and 01-03 and return in logging order.
	// Should it return errors? If so, what kind?
	logs := server.Search(?)
	require.Equal(t, logs, []string{
		`{"time":"2019-01-03","number":126}`,  // Notice that JSON format is compact
		`{"time":"2019-01-02","number":127}`,
		`{"time":"2019-01-01","number":128}`
	})

	// Find the record between 01-01 and 01-03, select 'number' only and return in logging order
	logs := server.Search(?)
	require.Equal(t, logs, []string{
		`{"number":126}`,
		`{"number":127}`,
		`{"number":128}`
	})
}