[![Travis CI](https://travis-ci.org/federico-arias/go-test.svg?branch=master)

## Instalación

Con Go modules:

```go
import "github.com/federico-arias/go-test"`
```

## Uso

Signatura de la función:

```go
func ParseTLV(tlv []byte) (map[string]string, error)
```

### Ejemplo

```go
r, err := tlv.ParseTLV([]byte("11A05AB398765UJ102N2300"))
```

Donde r es

```go
map[string]string{"05": "AB398765UJ1", "23": "00"}
```

Es decir, el valor indexado por el tipo.

## CI

https://travis-ci.org/federico-arias/go-test
