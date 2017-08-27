package steamtop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSteamTopGames(t *testing.T) {
	result, err := GetSteamTopGames()
	assert.NoError(t, err)
	assert.NotZero(t, len(result))

	t.Log(result[:10])
}
