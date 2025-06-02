## GENERATE PROTO MODEL
```bash
protoc --go_out=. --go-grpc_out=. ./helloworld.proto
```

### IF FOUND ERROR LIKE
<pre>protoc-gen-go: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go_out: protoc-gen-go: Plugin failed with status code 1.
</pre>

#### COPY AND RUN
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```
