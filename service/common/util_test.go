package common

import (
	"testing"
)

/*
 * go test ./service/common/
 */
func TestGenerateUID(t *testing.T) {
	uid := GenerateUID()
	if uid == "" {
		t.Error("UIDが生成されていない。")
	}
}
