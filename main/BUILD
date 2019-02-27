load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

exports_files(["run.sh"])

go_binary(
    name = "server",
    srcs = [
        "events.go",
        "main.go",
        "net.go",
    ],
    visibility = ["//visibility:public"],
)
