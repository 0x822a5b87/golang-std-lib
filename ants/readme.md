# ants

## Options

> ants 的源码中有一个和 java 区别很大的写法，这里记录一下
>
> `Options` 是一个结构体，包含了对 `ants` 协程池的配置；
>
> `Option` 是一个参数为 `*Options` 的函数；
>
> `loadOptions` 接受 `...Option` 作为参数并返回一个 `Options`，并对于数组中的每一个 `Option` 函数都执行一次。
>
> 所以，可以通过 `WithLogger` 这样来设置 `Logger`
>
> **这样的有点是，可以更自由的去定制我们的 Option 函数。**

```go
// Option represents the optional function.
type Option func(opts *Options)
```

```go
// Options contains all options which will be applied when instantiating an ants pool.
type Options struct {
	// ExpiryDuration is a period for the scavenger goroutine to clean up those expired workers,
	// the scavenger scans all workers every `ExpiryDuration` and clean up those workers that haven't been
	// used for more than `ExpiryDuration`.
	ExpiryDuration time.Duration

	// PreAlloc indicates whether to make memory pre-allocation when initializing Pool.
	PreAlloc bool

	// Max number of goroutine blocking on pool.Submit.
	// 0 (default value) means no such limit.
	MaxBlockingTasks int

	// When Nonblocking is true, Pool.Submit will never be blocked.
	// ErrPoolOverload will be returned when Pool.Submit cannot be done at once.
	// When Nonblocking is true, MaxBlockingTasks is inoperative.
	Nonblocking bool

	// PanicHandler is used to handle panics from each worker goroutine.
	// if nil, panics will be thrown out again from worker goroutines.
	PanicHandler func(interface{})

	// Logger is the customized logger for logging info, if it is not set,
	// default standard logger from log package is used.
	Logger Logger
}
```

```go
func loadOptions(options ...Option) *Options {
	opts := new(Options)
	for _, option := range options {
		option(opts)
	}
	return opts
}
```

```go
// WithLogger sets up a customized logger.
func WithLogger(logger Logger) Option {
	return func(opts *Options) {
		opts.Logger = logger
	}
}
```

