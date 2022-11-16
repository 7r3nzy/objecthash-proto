#!/bin/bash -eu
#
# Copyright 2017 The ObjectHash-Proto Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

readonly TEST_PROTOS_DIR="$(cd "$(dirname "$0")"; pwd)"
readonly SCHEMA_DIR="${TEST_PROTOS_DIR}/schema"
readonly GENERATED_DIR="${TEST_PROTOS_DIR}/generated"
readonly LATEST_COMMIT="${TEST_PROTOS_DIR}/latest_commit.txt"
readonly LATEST_DIR="${GENERATED_DIR}/latest"

echo $LATEST_DIR
readonly TMP_GOPATH="./tmp/go"
trap "rm -rf ${TMP_GOPATH}" EXIT
mkdir -p $TMP_GOPATH

readonly GOPATH="${TMP_GOPATH}"
readonly GOBIN="${TMP_GOPATH}/bin"
readonly PATH="${TMP_GOPATH}/bin:${PATH}"

readonly TMP_PROTOC_PATH="./tmp/protoc"
trap "rm -rf ${TMP_PROTOC_PATH}" EXIT
mkdir -p $TMP_PROTOC_PATH

readonly PROTOC_VERSION="21.9"
readonly PROTOC_URL="https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip"
readonly PROTOC_BIN="${TMP_PROTOC_PATH}/protoc/bin/protoc"

generate_protos() {
  local schema_dir="$1"
  local output_dir="$2"

  # Make sure the directory exists and is empty.
  rm -rf "${output_dir}"
  mkdir -p "${output_dir}"

  for version in proto2 proto3; do
    mkdir -p "${output_dir}/${version}"
    "${PROTOC_BIN}" \
      --proto_path="${schema_dir}/${version}" \
      --go_opt=paths=source_relative \
      --go_out="${output_dir}/${version}" \
      "${schema_dir}/${version}"/*.proto
  done
}

install_protoc() {
  curl --location "${PROTOC_URL}" \
    --output "${TMP_PROTOC_PATH}/protoc.zip"
  unzip "${TMP_PROTOC_PATH}/protoc.zip" \
    -d "${TMP_PROTOC_PATH}/protoc"
}

install_protoc_gen_go() {
  GO_BIN="$HOME/${TMP_GOPATH}/bin" go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
}

get_protoc_gen_commit() {
  (
    cd "${TMP_GOPATH}/src/google.golang.org/protobuf"
    git rev-parse HEAD
  )
}

run() {
  rm -rf "${GENERATED_DIR}"
  mkdir -p "${GENERATED_DIR}"

  install_protoc
  install_protoc_gen_go
  generate_protos "${SCHEMA_DIR}" "${LATEST_DIR}"

  #local commit="$(get_protoc_gen_commit)"
  #echo "${commit}" > "${LATEST_COMMIT}"
}

run
