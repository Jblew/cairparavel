#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"
set -e

echo "# Installing cairparavel"

echo "## Deploying functions"
./functions/deploy.sh
echo "## Deploying functions done"

echo "## Building web"
./web/build.sh
echo "## Building web done"

echo "## Deploying firestore and hosting"
firebase deploy --only firestore,hosting
echo "## Deploying firestore and hosting done"

echo "# Cairparavel installation done"
