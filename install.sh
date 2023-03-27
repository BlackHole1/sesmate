#!/bin/sh

# Set the base URL for the download
base_url="https://github.com/BlackHole1/sesmate/releases/latest/download/"

# Determine the OS and architecture
os=$(uname -s | tr '[:upper:]' '[:lower:]')
case $(uname -m) in
    x86_64)
        arch="amd64"
        ;;
    arm64*)
        arch="arm64"
        ;;
    *)
        echo "Unsupported architecture: $(uname -m)"
        exit 1
        ;;
esac

# Construct the URL for the binary based on the OS and architecture
url="${base_url}sesmate-${os}-${arch}"

# Download the binary
if ! curl -L "${url}" -o "/usr/local/bin/sesmate"; then
    echo "Failed to download sesmate from ${url}"
    exit 1
fi

# Make the binary executable
chmod +x "/usr/local/bin/sesmate"

echo "Successfully downloaded and installed sesmate to /usr/local/bin/sesmate"
exit 0
