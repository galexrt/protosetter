package testdata

import (
	"github.com/galexrt/protosetter/testdata/proto"
)

func testInvalid(t *proto.Test) {
	optBool := true
	t.D = 1.5                                  // want `avoid direct access to proto field t\.D, use t\.SetD\(1.5\) instead`
	t.S = "test"                               // want `avoid direct access to proto field t\.S, use t\.SetS\("test"\) instead`
	t.B = []byte{1, 2, 3}                      // want `avoid direct access to proto field t\.B, use t\.SetB\(\[\]byte\{1, 2, 3\}\) instead`
	t.I32 = 7                                  // want `avoid direct access to proto field t\.I32, use t\.SetI32\(7\) instead`
	t.I32 += 2                                 // want `avoid direct access to proto field t\.I32, use t\.SetI32\(t\.GetI32\(\) \+ 2\) instead`
	t.I32 -= 3                                 // want `avoid direct access to proto field t\.I32, use t\.SetI32\(t\.GetI32\(\) - 3\) instead`
	t.I32++                                    // want `avoid direct access to proto field t\.I32, use t\.SetI32\(t\.GetI32\(\) \+ 1\) instead`
	t.I32--                                    // want `avoid direct access to proto field t\.I32, use t\.SetI32\(t\.GetI32\(\) - 1\) instead`
	t.Embedded = &proto.Embedded{}             // want `avoid direct access to proto field t\.Embedded, use t\.SetEmbedded\(&proto\.Embedded\{\}\) instead`
	t.Embedded.S = "nested"                    // want `avoid direct access to proto field t\.Embedded\.S, use t\.GetEmbedded\(\)\.SetS\("nested"\) instead`
	t.Embedded.S += "!"                        // want `avoid direct access to proto field t\.Embedded\.S, use t\.GetEmbedded\(\)\.SetS\(t\.GetEmbedded\(\)\.GetS\(\) \+ "!"\) instead`
	t.Embedded.Embedded.S = "deep"             // want `avoid direct access to proto field t\.Embedded\.Embedded\.S, use t\.GetEmbedded\(\)\.GetEmbedded\(\)\.SetS\("deep"\) instead`
	t.RepeatedEmbeddeds[0].S = "list"          // want `avoid direct access to proto field t\.RepeatedEmbeddeds\[0\]\.S, use t\.GetRepeatedEmbeddeds\(\)\[0\]\.SetS\("list"\) instead`
	t.Map = map[string]string{"test": "value"} // want `avoid direct access to proto field t\.Map, use t\.SetMap\(map\[string\]string\{"test": "value"\}\) instead`
	t.OptBool = &optBool                       // want `avoid direct access to proto field t\.OptBool, use t\.SetOptBool\(optBool\) instead`
	t.OptBool = nil                            // want `avoid direct access to proto field t\.OptBool, use t\.ClearOptBool\(\) instead`
}

func testValid(t *proto.Test) {
	t.SetD(1.5)
	t.SetS("test")
	t.SetB([]byte{1, 2, 3})
	t.SetI32(7)
	t.SetI32(t.GetI32() + 2)
	t.SetI32(t.GetI32() - 3)
	t.SetI32(t.GetI32() + 1)
	t.SetI32(t.GetI32() - 1)
	t.SetEmbedded(&proto.Embedded{})
	t.GetEmbedded().SetS("nested")
	t.GetEmbedded().SetS(t.GetEmbedded().GetS() + "!")
	t.GetEmbedded().GetEmbedded().SetS("deep")
	t.GetRepeatedEmbeddeds()[0].SetS("list")
	t.SetMap(map[string]string{"test": "value"})
}
