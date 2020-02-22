package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

var flagtests = []struct {
	in   []byte
	out  map[string]string
	desc string
}{
	{
		[]byte("11A05AB398765UJ102N2300"),
		map[string]string{"05": "AB398765UJ1", "23": "00"},
		"cadena bien formada",
	},
	{
		[]byte("11A05A😀398765UJ102N2300"),
		map[string]string{"05": "A😀398765UJ1", "23": "00"},
		"cadena bien formada con un codepoint mayor a un octeto",
	},
	/*{
		[]byte("11A05AB398765UJ102N230"),
		map[string]string{"05": "AB398765UJ1", "23": "00"},
		"cadena mal formada",
	},
	*/
}

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	rc := m.Run()

	// rc 0 means we've passed,
	// and CoverMode will be non empty if run with -cover
	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < 0.8 {
			fmt.Println("Tests passed but coverage failed at", c)
			rc = -1
		}
	}
	os.Exit(rc)
}

func TestParseTLV(t *testing.T) {
	for _, ta := range flagtests {
		r, err := ParseTLV(ta.in)
		t.Logf(ta.desc)
		if err != nil {
			t.Errorf("error: %s", err.Error())
		}
		if !reflect.DeepEqual(r, ta.out) {
			t.Errorf("recieved %#v, expected %#v", r, ta.out)
		}
	}
}
