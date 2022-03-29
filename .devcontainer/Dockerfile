# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.192.0/containers/dotnet/.devcontainer/base.Dockerfile

# [Choice] .NET version: 5.0, 3.1, 2.1
ARG VARIANT="5.0"
FROM mcr.microsoft.com/vscode/devcontainers/dotnet:0-${VARIANT}

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# [Option] Install Azure CLI
ARG INSTALL_AZURE_CLI="false"
COPY library-scripts/*.sh library-scripts/*.env /tmp/library-scripts/
RUN if [ "$INSTALL_AZURE_CLI" = "true" ]; then bash /tmp/library-scripts/azcli-debian.sh; fi \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/library-scripts

# Add python3
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive && apt-get install -y software-properties-common gcc \
     && apt-get -y install --no-install-recommends python3.9 python3-doc python3-distutils python3-venv python3-pip python3-apt python3-dev \
     && python3 -m pip install --no-cache-dir --upgrade pip

# Install commitizen as a global node package
RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && echo '{ \"path\": \"cz-customizable\" }' > ~/.czrc && npm install -g commitizen cz-customizable" 2>&1