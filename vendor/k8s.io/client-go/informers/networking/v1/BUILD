load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "networking",
    srcs = [
        "ingress.go",
        "ingressclass.go",
        "interface.go",
        "networkpolicy.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/client-go/informers/networking/v1",
    importpath = "k8s.io/client-go/informers/networking/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/api/networking/v1:networking",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/runtime",
        "//vendor/k8s.io/apimachinery/pkg/watch",
        "//vendor/k8s.io/client-go/informers/internalinterfaces",
        "//vendor/k8s.io/client-go/kubernetes",
        "//vendor/k8s.io/client-go/listers/networking/v1:networking",
        "//vendor/k8s.io/client-go/tools/cache",
    ],
)
