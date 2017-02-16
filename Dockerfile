FROM alpine:3.5

COPY .dist/kube-distribution /usr/bin/kube-distribution

ENTRYPOINT ["/usr/bin/kube-distribution"]