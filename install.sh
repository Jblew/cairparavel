#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"

echo "# Installing cairparavel"

echo "## Deploying firestore and hosting"
firebase deploy --only firestore,hosting
echo "## Deploying firestore and hosting done"

echo "# Cairparavel installation done"
