package http

import (
	stdhttp "net/http"
	"strconv"

	"github.com/sibyakin/indenter"
)

type JSON struct {
	i indenter.JSON
	l int
}

func NewJSON() JSON {
	result := JSON{}

	return result
}

func (h JSON) Run(hostport string, indentLen int) {
	h.l = indentLen
	stdhttp.ListenAndServe(hostport, h)
}

func (h JSON) ServeHTTP(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	h.i = indenter.NewJSONUnbuffered(w)
	h.i.SetInput(r.Body)

	i, err := strconv.Atoi(r.URL.Query().Get("indent"))
	if err == nil {
		h.l = i
	}
	h.i.SetIndent(h.l)

	h.i.Run()
}
