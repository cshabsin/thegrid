
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# gazelle:repository_macro go_repositories.bzl%go_repositories
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "07fe7474d1b93083a744c903461d36d013165278013497e47e42316237e05aa6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.45.0/rules_go-v0.45.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.45.0/rules_go-v0.45.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "1a4c4431a4b2e92595e6855e433c66a73e718689b2f52e401e9b1860f8310b88",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.zip",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.22.5")

gazelle_dependencies()

