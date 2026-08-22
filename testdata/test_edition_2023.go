package testdata

import (
	"github.com/galexrt/protosetter/testdata/proto"
)

func testValidEdition2023(t *proto.TestEdition2023) {
	t.SetValue(map[string]string{"test": "test"})
	_ = t.GetValue()
}
