load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "flag",
    srcs = [
        "ciphersuites_flag.go",
        "colon_separated_multimap_string_string.go",
        "configuration_map.go",
        "flags.go",
        "langle_separated_map_string_string.go",
        "map_string_bool.go",
        "map_string_string.go",
        "namedcertkey_flag.go",
        "noop.go",
        "omitempty.go",
        "sectioned.go",
        "string_flag.go",
        "string_slice_flag.go",
        "tristate.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/component-base/cli/flag",
    importpath = "k8s.io/component-base/cli/flag",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/spf13/cobra",
        "//vendor/github.com/spf13/pflag",
        "//vendor/k8s.io/apimachinery/pkg/util/sets",
        "//vendor/k8s.io/klog/v2:klog",
    ],
)
