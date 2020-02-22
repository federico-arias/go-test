package main

import (
	"reflect"
	"testing"
)

var flagtests = []struct {
	in   []byte
	out  map[string]string
	desc string
}{
	/*{
		[]byte("11A05A😀398765UJ102N2300"),
		map[string]string{"05": "A😀398765UJ1", "23": "00"},
		"cadena con un codepoint mayor a un octeto",
	},*/
	{
		[]byte("11A05AB398765UJ102N2300"),
		map[string]string{"05": "AB398765UJ1", "23": "00"},
		"cadena correcta",
	},
}

func TestParseTLV(t *testing.T) {
	for _, ta := range flagtests {
		r, err := ParseTLV(ta.in)
		t.Log(ta.desc)
		if err != nil {
			t.Errorf("error: %s", err.Error())
		}
		if reflect.DeepEqual(r, ta.out) {
			t.Errorf("expected %#v, recieved %#v", r, ta.out)
		}
	}
}
