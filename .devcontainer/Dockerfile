# Official Go image (Debian-based). Pin a version, e.g. golang:1.27, for reproducibility.
FROM golang:1

ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Non-root user with sudo (git is already included in the base image)
RUN apt-get update \
 && apt-get install -y --no-install-recommends sudo \
 && rm -rf /var/lib/apt/lists/* \
 && groupadd --gid $USER_GID $USERNAME \
 && useradd --uid $USER_UID --gid $USER_GID -m -s /bin/bash $USERNAME \
 && echo "$USERNAME ALL=(root) NOPASSWD:ALL" > /etc/sudoers.d/$USERNAME

# Go tooling: language server, debugger, linter
ENV GOBIN=/usr/local/bin
RUN go install golang.org/x/tools/gopls@latest \
 && go install github.com/go-delve/delve/cmd/dlv@latest \
 && go install honnef.co/go/tools/cmd/staticcheck@latest \
 && go install mvdan.cc/gofumpt@latest \
 && go clean -cache -modcache \
 && mkdir -p /go/pkg/mod \
 && chown -R $USERNAME:$USERNAME /go

USER $USERNAME