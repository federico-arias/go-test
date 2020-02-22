package main

import (
	"bytes"
	"errors"
	"io"
	"strconv"
)

type TLVReader struct {
	reader *bytes.Reader
	err    error
	t      string
	l      int64
	isAlfa bool
}

var (
	ErrTipoInvalido = errors.New("tipo inválido: no cumple largo")
	//ErrTipoDesconocido = errors.New("tipo inválido: solo se acepta A o N")
)

// Prepara el escaneo del valor del TLV
// Asume que estarán dentro de los valores ASCII.
func (r *TLVReader) Next() bool {
	// Verifica EOF
	if r.reader.Len() < 1 {
		return false
	}
	// Verifica que queden bytes suficientes
	if r.reader.Len() < 5 {
		r.err = ErrTipoInvalido
		return false
	}
	// Lee el largo
	largo := make([]byte, 2)
	_, err := r.reader.Read(largo)
	if err != nil {
		r.err = err
		return false
	}
	r.l, err = strconv.ParseInt(string(largo), 10, 0)
	if err != nil {
		r.err = err
		return false
	}
	// Lee el tipo
	tipo0 := make([]byte, 1)
	_, err = r.reader.Read(tipo0)
	if err != nil {
		r.err = err
		return false
	}
	// Lee si es alfanumérico
	alfa, err := checkAlfa(tipo0)
	if err != nil {
		r.err = err
		return false
	}
	r.isAlfa = alfa

	tipo1 := make([]byte, 2)
	_, err = r.reader.Read(tipo1)
	if err != nil {
		r.err = err
		return false
	}
	r.t = string(tipo1)
	// Guarda el tipo como alfanumérico
	return true
}

// Lee un byte a la vez. Como no siempre se cumple la
// equivalencia entre 1 octeto == 1 caracter, detecta
// esta discrepancia y modifica las runas a leer.
//
// Asume codificación UTF-8.
func (r *TLVReader) Scan(m map[string]string) (err error) {
	var ch rune
	for i := 0; i < int(r.l); i++ {
		ch, _, err = r.reader.ReadRune()
		if err == io.EOF {
			r.err = errors.New("segmento valor de largo insuficiente")
			return r.err
		}
		if err != nil {
			return err
		}
		m[r.t] = m[r.t] + string(ch)
	}
	return nil
}

func (r *TLVReader) Err() error {
	return r.err
}

func ParseTLV(tlv []byte) (map[string]string, error) {
	ret := make(map[string]string)
	r := &TLVReader{}
	r.reader = bytes.NewReader(tlv)

	for r.Next() {
		err := r.Scan(ret)
		if err != nil {
			return ret, err
		}
	}

	if err := r.Err(); err != nil {
		return ret, err
	}
	return ret, nil

}

func checkAlfa(b []byte) (bool, error) {
	switch string(b) {
	case "A":
		return true, nil
	case "N":
		return false, nil
	default:
		return false, errors.New("tipo inválido")
	}
}
