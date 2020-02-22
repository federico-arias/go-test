## Instalación

Con Go modules:

```
import "github.com/federico-arias/go-test"`
```

## Uso

Signatura de la función:

```
func ParseTLV(tlv []byte) (map[string]string, error)
```

### Ejemplo

```
r, err := go-test.ParseTLV(byte[]("11A05AB398765UJ102N2300"))
```

Donde r es

```
map[string]string{"05": "AB398765UJ1", "23": "00"}
```

Es decir, el valor indexado por el tipo.
