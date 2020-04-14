FROM registry.access.redhat.com/ubi8-minimal
COPY bin/grpc-demo-order /
EXPOSE 8080 5000
CMD ["/grpc-demo-order"]