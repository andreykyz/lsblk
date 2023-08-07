package lsblk

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed lvs.json
var JSONTest []byte

func TestLVS(t *testing.T) {
	lvsRsp := &Report{}
	err := json.Unmarshal(JSONTest, &lvsRsp)
	assert.NoError(t, err)
}
