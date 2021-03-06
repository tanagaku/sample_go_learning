package trace

import (
	"fmt"
	"io"
)

//Tracerはコード内での出来事を記録できるオブジェクトを返すインターフェースです
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

//OffはTraceメソッドの呼び出しを無視するTracerをかえします
func Off() Tracer {
	return &nilTracer{}
}
