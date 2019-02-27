def _dev_tcp_impl(ctx):
    ctx.actions.expand_template(
        template = ctx.file._template,
        substitutions = {
            "%{SRC_ADDR}": ctx.attr.listen_addr,
            "%{DST_ADDR}": ctx.attr.target_addr,
            "%{BIN}": ctx.file._server.short_path,
        },
        is_executable = True,
        output = ctx.outputs.bin,
    )

    return [DefaultInfo(executable = ctx.outputs.bin, runfiles = ctx.runfiles(files = [ctx.file._server]))]

dev_tcp = rule(
    implementation = _dev_tcp_impl,
    attrs = {
        "listen_addr": attr.string(),
        "target_addr": attr.string(),
        "_server": attr.label(default = "@bazel_dev_tcp//:server", allow_single_file = True),
        "_template": attr.label(default = "@bazel_dev_tcp//:run.sh", allow_single_file = True),
    },
    executable = True,
    outputs = {
        "bin": "%{name}-bin"
    }
)
