#!/usr/bin/env bash

# Copyright 2023 The KCP Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Usage: $0 --create-folder --verbose

# This script installs KCP-Edge binaries to a folder of choice
#
# Arguments:
# [--version release_version] set a specific KCP-Edge release version, default: latest
# [--os linux|darwin] set a specific OS type, default: autodetect
# [--arch amd64|arm64] set a specific architecture type, default: autodetect
# [--folder installation_folder] sets the installation folder, default: $PWD/kcp-edge
# [--create-folder] create the instllation folder, if it does not exist
# [-V|--verbose] verbose output

set -e

kcp_edge_version=""
kcp_edge_os=""
kcp_edge_arch=""
kcp_edge_folder=""
kcp_edge_create_folder="false"
verbose="false"

get_os_type() {
  case "$OSTYPE" in
      linux*)   echo "linux" ;;
      darwin*)  echo "darwin" ;;
      *)        echo "Unsupported operating system type: $OSTYPE" >&2 ; exit 1 ;;
  esac
}

get_arch_type() {
  case "$HOSTTYPE" in
      x86_64*)    echo "amd64" ;;
      aarch64*) echo "arm64" ;;
      *)        echo "Unsupported architecture type: $HOSTTYPE" >&2 ; exit 1 ;;
  esac
}

get_latest_version() {
    curl -sL https://github.com/kcp-dev/edge-mc/releases/latest | grep "</h1>" | head -n 1 | sed -e 's/<[^>]*>//g' | xargs
}

get_full_path() {
    echo "$(cd "$1"; pwd)"
}

while (( $# > 0 )); do
    case "$1" in
    (--version)
        if (( $# > 1 ));
        then { kcp_edge_version="$2"; shift; }
        else { echo "$0: missing release version" >&2; exit 1; }
        fi;;
    (--os)
        if (( $# > 1 ));
        then { kcp_edge_os="$2"; shift; }
        else { echo "$0: missing OS type" >&2; exit 1; }
        fi;;
    (--arch)
        if (( $# > 1 ));
        then { kcp_edge_arch="$2"; shift; }
        else { echo "$0: missing architecture type" >&2; exit 1; }
        fi;;
    (--folder)
        if (( $# > 1 ));
        then { kcp_edge_folder="$2"; shift; }
        else { echo "$0: missing installation folder" >&2; exit 1; }
        fi;;
    (--create-folder)
        kcp_edge_create_folder="true";;
    (--verbose|-V)
        verbose="true";;
    (-h|--help)
        echo "Usage: $0 [--version release_version] [--os linux|darwin] [--arch amd64|arm64] [--folder installation_folder] [--create-folder] [-V|--verbose]"
        exit 0;;
    (-*)
        echo "$0: unknown flag" >&2 ; exit 1;
        exit 1;;
    (*)
        echo "$0: unknown positional argument" >&2; exit 1;
        exit 1;;
    esac
    shift
done

if [ "$kcp_edge_version" == "" ]; then
    kcp_edge_version=$(get_latest_version)
fi
if [ "$kcp_edge_os" == "" ]; then
    kcp_edge_os=$(get_os_type)
fi
if [ "$kcp_edge_arch" == "" ]; then
    kcp_edge_arch=$(get_arch_type)
fi
if [ "$kcp_edge_folder" == "" ]; then
    kcp_edge_folder="$PWD/kcp-edge"
fi

if [ -d "$kcp_edge_folder" ]; then :
elif [ "$kcp_edge_create_folder" == "true" ]; then
    if [ $verbose == "true" ] ; then
        echo "Creating folder: $kcp_edge_folder"
    fi
    mkdir -p "$kcp_edge_folder"
else
    echo "Specified folder does not exist: $kcp_edge_folder" >&2; exit 1;
fi

if [ $verbose == "true" ]
then
    echo "Downloading KCP-Edge $kcp_edge_version $kcp_edge_os/$kcp_edge_arch..."
    curl -SL -o kcp-edge.tar.gz "https://github.com/kcp-dev/edge-mc/releases/download/${kcp_edge_version}/kcp-edge_${kcp_edge_version}_${kcp_edge_os}_$kcp_edge_arch.tar.gz"
else
    curl -sSL -o kcp-edge.tar.gz "https://github.com/kcp-dev/edge-mc/releases/download/${kcp_edge_version}/kcp-edge_${kcp_edge_version}_${kcp_edge_os}_$kcp_edge_arch.tar.gz"
fi

if [ $verbose == "true" ]
then
    echo "Extracting archive to: $kcp_edge_folder"
fi
tar -zxf kcp-edge.tar.gz -C $kcp_edge_folder

if [ $verbose == "true" ]
then
    echo "Cleaning up..."
fi
rm kcp-edge.tar.gz

if [[ ! ":$PATH:" == *":$(get_full_path $kcp_edge_folder/bin):"* ]]; then
    echo "Add KCP-Edge folder to your path: export PATH="\$PATH:$(get_full_path $kcp_edge_folder/bin)""
fi