load("@bazel_dev_tcp//:rules.bzl", "dev_tcp")

java_binary(
    name = "echo",
    main_class = "EchoServer",
    srcs = ["EchoServer.java"],
    jvm_flags =
        select({
            "//:dev": ["-Decho-server.port=18000"],
            "//conditions:default": ["-Decho-server.port=8000"]
        })
)

dev_tcp(
    name = "echo_dev",
    listen_addr = "localhost:8000",
    tags = [
        "ibazel_notify_changes"
    ],
    target_addr = "localhost:18000",
)
