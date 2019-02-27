# bazel-dev-tcp

An addition to ibazel that helps developing TCP services.

## Example

Consider a Java TCP server at port 8000.

```
java_binary(
    name = "echo",
    main_class = "EchoServer",
    srcs = ["EchoServer.java"],
    jvm_flags = ["-Decho-server.port=8000"],
)
```

Add `dev_tcp` to have a server running at 8000 that proxies the Java service at port 18000. From the time that ibazel detects a source change to when ibazel restarts the service, the proxy server will hold all new TCP connections.

```
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
