# bazel-dev-tcp

An addition to ibazel that helps developing TCP services.

## Getting started

Add to WORKSPACE

```python
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.17.1/rules_go-0.17.1.tar.gz"],
    sha256 = "6776d68ebb897625dead17ae510eac3d5f6342367327875210df44dbe2aeeb19",
)
load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()

http_archive(
    name = "bazel_dev_tcp",
    urls = ["https://github.com/pauldraper/bazel-dev-tcp/archive/ca4eefa9fc1ae3b8cbb6642fa429e3a2e4ceadf1.zip"],
    sha256 = "f6dd7c36a1de1554175abde27e3805198668958b974d90512498532795d848d2",
    strip_prefix = "bazel-dev-tcp-ca4eefa9fc1ae3b8cbb6642fa429e3a2e4ceadf1/main"
)
```


## Example

Consider a Java TCP server at port 8000.

```python
java_binary(
    name = "echo",
    main_class = "EchoServer",
    srcs = ["EchoServer.java"],
    jvm_flags = ["-Decho-server.port=8000"],
)
```

Add `dev_tcp` to have a server running at 8000 that proxies the Java service at port 18000. From the time that ibazel detects a source change to when ibazel restarts the service, the proxy server will hold all new TCP connections.

```python
load("@bazel_dev_tcp//:rules.bzl", "dev_tcp")

java_binary(
    name = "echo",
    main_class = "EchoServer",
    srcs = ["EchoServer.java"],
    jvm_flags = ["-Decho-server.port=8000"],
)

dev_tcp(
    name = "echo_dev",
    listen_addr = "localhost:8000",
    tags = [
        "ibazel_notify_changes"
    ],
    target_addr = "localhost:18000",
)
```

You can see an example in [tests/java](test/java) including a optionally enabling this behavior. (run vs run-dev)
