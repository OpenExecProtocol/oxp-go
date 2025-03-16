# Health

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Tools

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go">oxp</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go#ToolListResponse">ToolListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go">oxp</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go#ToolCallResponse">ToolCallResponse</a>

Methods:

- <code title="get /tools">client.Tools.<a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go#ToolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go">oxp</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go#ToolListParams">ToolListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go">oxp</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go#ToolListResponse">ToolListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tools/call">client.Tools.<a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go#ToolService.Call">Call</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go">oxp</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go#ToolCallParams">ToolCallParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go">oxp</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/oxp-go#ToolCallResponse">ToolCallResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
