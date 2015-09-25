package writer

import (
	"fmt"
	"io"

	"github.com/munnerz/haproxy/api"
)

func WriteString(w io.Writer, s string) (int, error) {
	return w.Write([]byte(fmt.Sprintf("%s\n", s)))
}

func WriteWriteable(w io.Writer, o api.Writeable) (int, error) {
	t := 0

	for _, s := range o.Strings() {
		n, err := WriteString(w, s)
		t += n

		if err != nil {
			return t, fmt.Errorf("Error writing writeable: %s", err.Error())
		}
	}

	return t, nil
}
