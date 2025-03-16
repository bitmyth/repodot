Dockerfile 优化建议

使用更小的基础镜像：

如果不需要 alpine 的额外功能，可以使用 scratch 镜像（完全空的基础镜像）：

FROM scratch AS runtime
COPY --from=builder /bin/myapp /app/myapp
CMD ["/app/myapp"]


如果需要支持多种 CPU 架构（如 ARM、x86），可以使用 docker buildx 构建多平台镜像：

