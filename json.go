package indenter

import (
	"encoding/json"
	"io"
)

const jsonext = "json"

type JSON struct {
	core
}

// NewJSON is the default and recommended json indenter.
// You can use any io.Writer to write output but
// STDOUT will be used as default if you use nil.
// You can pass optional slice of filenames to
// work with. Otherwise it will pick up json files
// from CWD by itself.
func NewJSON(w io.Writer, filenames ...string) JSON {
	return JSON{core: newCore(true, jsonext, w, filenames...)}
}

// NewJSONUnbuffered is the same as NewJSON but
// without any internal buffering. Performance
// impact may be significant.
func NewJSONUnbuffered(w io.Writer, filenames ...string) JSON {
	return JSON{core: newCore(false, jsonext, w, filenames...)}
}

func (j JSON) Run() {
	if j.buffered {
		j.runBuffered()
	} else {
		j.runUnbuffered()
	}

	for _, v := range j.onshutdown {
		v()
	}
}

func (j JSON) runBuffered() {
	defer j.bufout.Flush()

	for _, v := range j.in {
		j.bufin.Reset(v)
		r, err := io.ReadAll(&j.bufin)
		if err == nil {
			json.Indent(&j.b, r, "", j.indent)
			j.bufout.Write(j.b.Bytes())
		}
		j.b.Reset()
	}

}

func (j JSON) runUnbuffered() {
	for _, v := range j.in {
		r, err := io.ReadAll(v)
		if err == nil {
			json.Indent(&j.b, r, "", j.indent)
			j.out.Write(j.b.Bytes())
		}
		j.b.Reset()
	}
}
