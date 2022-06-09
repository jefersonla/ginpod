#!/bin/bash

export GOOGLE_CLOUD_VERSION=389.0.0
export GOOGLE_CLOUD_ARCH=x86_64
export GOOGLE_CLOUD_INSTALLER_NAME=gcp_installer.tar.gz

# Install Google Cloud
mkdir -p ~/Apps/gcp
    cd ~/Apps/gcp && \
    curl -o "$GOOGLE_CLOUD_INSTALLER_NAME" \
        "https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-$GOOGLE_CLOUD_VERSION-linux-$GOOGLE_CLOUD_ARCH.tar.gz" && \
    tar -zxvf "$GOOGLE_CLOUD_INSTALLER_NAME" && \
    ./google-cloud-sdk/install.sh -q && \
    echo "source /home/gitpod/Apps/gcp/google-cloud-sdk/path.bash.inc" >> ~/.profile &&
    source /home/gitpod/Apps/gcp/google-cloud-sdk/path.bash.inc 