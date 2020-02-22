package tlv

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
	err  error
}{
	{
		[]byte("11A05AB398765UJ102N2300"),
		map[string]string{"05": "AB398765UJ1", "23": "00"},
		"lee una cadena bien formada",
		nil,
	},
	{
		[]byte("11A05A😀398765UJ102N2300"),
		map[string]string{"05": "A😀398765UJ1", "23": "00"},
		`lee una cadena bien formada con un codepoint mayor
		a un octeto`,
		nil,
	},
	{
		[]byte("11C05AB398765UJ102N2300"),
		map[string]string{"05": "AB398765UJ1", "23": "00"},
		"reconce error de tipo de valor",
		ErrTipoDesconocido,
	},
	{
		[]byte(""),
		nil,
		"reconocer error de campo vacío",
		ErrCadenaVacia,
	},
	{
		[]byte("11A05AB398765UJ102N"),
		nil,
		"reconce error de cadena no terminada",
		ErrTipoInvalido,
	},
	{
		[]byte("11N05A😀398765UJ102N2300"),
		nil,
		`reconce error de concordancia tipo/valor`,
		ErrValorInvalido,
	},
}

func TestMain(m *testing.M) {
	rc := m.Run()

	if rc == 0 && testing.CoverMode() != "" {
		c := testing.Coverage()
		if c < 0.8 {
			fmt.Println(
				"Los test pasan pero falla el coverage en",
				c,
			)
			rc = -1
		}
	}
	os.Exit(rc)
}

func TestParseTLV(t *testing.T) {
	for _, ta := range flagtests {
		r, err := ParseTLV(ta.in)
		t.Logf(ta.desc)
		if err != nil && ta.err == nil {
			t.Errorf("error: %s", err.Error())
		}
		if err != ta.err {
			t.Errorf(
				"ParseTLV(%s) = _, %#v \n, expected %#v",
				ta.in,
				err,
				ta.err,
			)
			continue
		}
		if ta.err == nil && !reflect.DeepEqual(r, ta.out) {
			t.Errorf("ParseTLV(%s) = %#v, expected %#v", ta.in, r, ta.out)
		}
	}
}
