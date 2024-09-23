package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tracker-tv/actor-api/testutils"
)

func TestGetActors_Success(t *testing.T) {
	cfg, err := testutils.SetupConfig()
	assert.NoError(t, err)

	db, err := testutils.SetupDB(cfg.DB)
	assert.NoError(t, err)

	model := PgActorModel{DB: db}

	actors, err := model.GetActors()
	assert.NoError(t, err)

	assert.Len(t, actors, 3)
}
