#!/usr/bin/env bash
WEB_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
set -e

npm --prefix "${WEB_DIR}" install
npm --prefix "${WEB_DIR}" run build
