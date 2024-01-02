ARG VARIANT=bullseye
FROM mcr.microsoft.com/vscode/devcontainers/go:1.21-${VARIANT}

ARG HOME=/home/vscode
ARG UID=1000
ARG GID=${UID}

RUN install -d -m 0755 -o vscode -g ${GID} ${HOME}/.config
COPY --chmod=0755 --chown=${UID}:${GID} config/atuin.toml ${HOME}/.config/atuin.toml
COPY --chmod=0755 --chown=${UID}:${GID} config/starship.toml ${HOME}/.config/starship.toml

# Add scripts to run during final layer additions
COPY --chmod=0755 --chown=${UID}:${GID} scripts/* /tmp/install/

ENV PATH=${HOME}/.local/bin:${PATH}