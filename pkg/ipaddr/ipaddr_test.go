package ipaddr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApiAddressQuery(t *testing.T) {
	ip := "123.139.81.219"
	city, err := IPAddressQuery(ip)
	require.NoError(t, err)
	require.Equal(t, city, "西安")
}
