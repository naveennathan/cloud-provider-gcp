load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "storage",
    srcs = [
        "api_object_versioner.go",
        "continue.go",
        "doc.go",
        "errors.go",
        "interfaces.go",
        "selection_predicate.go",
        "util.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/apiserver/pkg/storage",
    importpath = "k8s.io/apiserver/pkg/storage",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/api/errors",
        "//vendor/k8s.io/apimachinery/pkg/api/meta",
        "//vendor/k8s.io/apimachinery/pkg/api/validation/path",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/fields",
        "//vendor/k8s.io/apimachinery/pkg/labels",
        "//vendor/k8s.io/apimachinery/pkg/runtime",
        "//vendor/k8s.io/apimachinery/pkg/types",
        "//vendor/k8s.io/apimachinery/pkg/util/validation/field",
        "//vendor/k8s.io/apimachinery/pkg/watch",
        "//vendor/k8s.io/apiserver/pkg/endpoints/request",
    ],
)
