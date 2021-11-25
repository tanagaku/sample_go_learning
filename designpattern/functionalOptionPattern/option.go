package option

import "log"

type options struct {
	param  string
	flag   bool
	logger *log.Logger
}

type Option interface {
	apply(*options)
}

type flagOption bool

func (f flagOption) apply(opts *options) {
	opts.flag = bool(f)
}

func WithFlag(b bool) Option {
	return flagOption(b) //flagOptiionに型変換
}

type paramOption string

func (p paramOption) apply(opts *options) {
	opts.param = string(p)
}

func WithParam(s string) Option {
	return paramOption(s)
}

type loggerOption struct {
	Log *log.Logger
}

func (l loggerOption) apply(opts *options) {
	opts.logger = l.Log
}

func WithLogger(log *log.Logger) Option {
	return loggerOption{Log: log}
}

//引数が可変なのがミソ
func Open(opts ...Option) (interface{}, error) {
	options := options{}

	for _, o := range opts {
		o.apply(&options)
	}

	return options, nil
}
