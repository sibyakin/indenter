package indenter

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var ErrIndentInvalid = errors.New("indenter: indent invalid")

type core struct {
	indent     string
	buffered   bool
	b          bytes.Buffer
	in         []io.Reader
	bufin      bufio.Reader
	out        io.Writer
	bufout     bufio.Writer
	onshutdown []func() error
}

func newCore(buffered bool, ext string, w io.Writer, filenames ...string) core {
	result := core{}

	result.indent = ""
	result.buffered = buffered
	result.in = make([]io.Reader, 0)
	result.onshutdown = make([]func() error, 0)

	result.walk(filenames)
	if len(result.in) == 0 {
		files, _ := filepath.Glob("./*." + ext)
		result.walk(files)
	}

	if result.buffered {
		result.bufout = *bufio.NewWriter(os.Stdout)
		if w != nil {
			result.bufout = *bufio.NewWriter(w)
		}
	} else {
		result.out = os.Stdout
		if w != nil {
			result.out = w
		}
	}

	return result
}

func (c *core) walk(filenames []string) {
	for _, v := range filenames {
		f, err := os.Open(v)
		if err == nil {
			c.in = append(c.in, f)
			c.onshutdown = append(c.onshutdown, f.Close)
		}
	}
}

func (c *core) SetIndent(indent int) error {
	if indent >= 0 && indent < 100 {
		c.indent = strings.Repeat(" ", indent)
		return nil
	}

	return ErrIndentInvalid
}

func (c *core) SetInput(r io.Reader) {
	if c.buffered {
		c.bufin = *bufio.NewReader(r)
	} else {
		c.in = []io.Reader{r}
	}
}
