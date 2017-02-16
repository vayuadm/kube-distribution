FROM alpine:3.5

COPY .dist/kube-distribution /usr/bin/kube-distribution

EXPOSE 5050

ENTRYPOINT ["/usr/bin/kube-distribution"]