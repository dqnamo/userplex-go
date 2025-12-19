# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/dqnamo/userplex-go">userplex</a>.<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go#UserIdentifyResponse">UserIdentifyResponse</a>

Methods:

- <code title="post /api/identify">client.Users.<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go#UserService.Identify">Identify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dqnamo/userplex-go">userplex</a>.<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go#UserIdentifyParams">UserIdentifyParams</a>) (<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go">userplex</a>.<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go#UserIdentifyResponse">UserIdentifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Logs

Response Types:

- <a href="https://pkg.go.dev/github.com/dqnamo/userplex-go">userplex</a>.<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go#LogBatchResponse">LogBatchResponse</a>
- <a href="https://pkg.go.dev/github.com/dqnamo/userplex-go">userplex</a>.<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go#LogNewResponse">LogNewResponse</a>

Methods:

- <code title="post /api/logs/batch">client.Logs.<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go#LogService.Batch">Batch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dqnamo/userplex-go">userplex</a>.<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go#LogBatchParams">LogBatchParams</a>) (<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go">userplex</a>.<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go#LogBatchResponse">LogBatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/log">client.Logs.<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go#LogService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/dqnamo/userplex-go">userplex</a>.<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go#LogNewParams">LogNewParams</a>) (<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go">userplex</a>.<a href="https://pkg.go.dev/github.com/dqnamo/userplex-go#LogNewResponse">LogNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
