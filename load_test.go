package conf

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testConf struct {
	Key string `json:"key"`
}

func TestLoadFileFrom(t *testing.T) {
	c := testConf{}
	err := LoadFileFrom("testdata", "conf1", &c)
	require.NoError(t, err)
	require.Equal(t, "value1", c.Key)
	err = LoadFileFrom("testdata", "conf2", &c)
	require.NoError(t, err)
	require.Equal(t, "value2", c.Key)
	err = LoadFileFrom("testdata", "conf3", &c)
	require.Error(t, err)
}