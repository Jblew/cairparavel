#!/usr/bin/env bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "${DIR}"
set -e

source ./functions.config.sh

if [ -z "${GCP_PROJECT_REGION}" ]; then
    echo "GCP_PROJECT_REGION env is not set"
    exit 1
fi

if [ -z "${GCP_PROJECT_ID}" ]; then
    echo "GCP_PROJECT_ID env is not set"
    exit 1
fi

./generate-config.sh

gcloud config set project "${GCP_PROJECT_ID}"

(
  echo "Deploying FnOnFirebaseAccountCreated"
  gcloud functions deploy FnOnFirebaseAccountCreated \
    --trigger-event providers/firebase.auth/eventTypes/user.create \
    --trigger-resource "${GCP_PROJECT_ID}" \
    --region "${GCP_PROJECT_REGION}" \
    --runtime go113 \
    --quiet \
    --memory "1024MB"
) &

(
  echo "Deploying FnOnNotificationQueued"
  gcloud functions deploy FnOnNotificationQueued \
    --trigger-event providers/cloud.firestore/eventTypes/document.create \
    --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/notifications/{uid}/messenger_queue/{notificationId}" \
    --region "${GCP_PROJECT_REGION}" \
    --runtime go113 \
    --quiet \
    --memory "256MB"
) &

(
  echo "Deploying FnOnCommentAdded"
  gcloud functions deploy FnOnCommentAdded \
    --trigger-event providers/cloud.firestore/eventTypes/document.create \
    --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/event_comments/{eventId}/messages/{commentId}" \
    --region "${GCP_PROJECT_REGION}" \
    --runtime go113 \
    --quiet \
    --memory "256MB"
) &

(
  echo "Deploying FnOnEventCreated"
  gcloud functions deploy FnOnEventCreated \
    --trigger-event providers/cloud.firestore/eventTypes/document.create \
    --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/events/{eventId}" \
    --region "${GCP_PROJECT_REGION}" \
    --runtime go113 \
    --quiet \
    --memory "256MB"
) &

(
  echo "Deploying FnOnEventModified"
  gcloud functions deploy FnOnEventModified \
    --trigger-event providers/cloud.firestore/eventTypes/document.update \
    --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/events/{eventId}" \
    --region "${GCP_PROJECT_REGION}" \
    --runtime go113 \
    --quiet \
    --memory "256MB"
) &

(
  echo "Deploying FnOnEventMemberSignupCreated"
  gcloud functions deploy FnOnEventMemberSignupCreated \
    --trigger-event providers/cloud.firestore/eventTypes/document.create \
    --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/events/{eventId}/signedMembers/{uid}" \
    --region "${GCP_PROJECT_REGION}" \
    --runtime go113 \
    --quiet \
    --memory "256MB"
) &

(
  echo "Deploying FnOnEventMemberSignupDeleted"
  gcloud functions deploy FnOnEventMemberSignupDeleted \
    --trigger-event providers/cloud.firestore/eventTypes/document.delete \
    --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/events/{eventId}/signedMembers/{uid}" \
    --region "${GCP_PROJECT_REGION}" \
    --runtime go113 \
    --quiet \
    --memory "256MB"
) &

(
  echo "Deploying FnOnEventVoteCreated"
  gcloud functions deploy FnOnEventVoteCreated \
    --trigger-event providers/cloud.firestore/eventTypes/document.create \
    --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/events/{eventId}/votes/{uid}" \
    --region "${GCP_PROJECT_REGION}" \
    --runtime go113 \
    --quiet \
    --memory "256MB"
) &

(
  echo "Deploying FnOnEventVoteModified"
  gcloud functions deploy FnOnEventVoteModified \
    --trigger-event providers/cloud.firestore/eventTypes/document.update \
    --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/events/{eventId}/votes/{uid}" \
    --region "${GCP_PROJECT_REGION}" \
    --runtime go113 \
    --quiet \
    --memory "256MB"
) &

(
echo "Deploying FnOnEventVoteDeleted"
gcloud functions deploy FnOnEventVoteDeleted \
  --trigger-event providers/cloud.firestore/eventTypes/document.delete \
  --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/events/{eventId}/votes/{uid}" \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --quiet \
  --memory "256MB"
) &


(
  echo "Deploying FnOnCronHandleEvents"
  CHECK_EVENTS_PUBSUB_TOPIC="cron-check-events-state"
  CHECK_EVENTS_PUBSUB_SUBSCRIPTION="sub_check-events-state"
  CHECK_EVENTS_SCHEDULER_JOB="job_check-events-state"
  gcloud pubsub topics create "${CHECK_EVENTS_PUBSUB_TOPIC}" || echo "Cannot create topic"
  gcloud pubsub subscriptions create "${CHECK_EVENTS_PUBSUB_SUBSCRIPTION}" --topic "${CHECK_EVENTS_PUBSUB_TOPIC}" || echo "Cannot create subscription"
  gcloud alpha scheduler jobs delete "${CHECK_EVENTS_SCHEDULER_JOB}" || echo "Failed to delete scheduler job"
  gcloud alpha scheduler jobs create pubsub "${CHECK_EVENTS_SCHEDULER_JOB}" \
    --topic "${CHECK_EVENTS_PUBSUB_TOPIC}" \
    --schedule "*/10 * * * *" \
    --message-body "SCHEDULE" || echo "Cannot recreate cron job"
  gcloud functions deploy FnOnCronHandleEvents \
    --trigger-topic "${CHECK_EVENTS_PUBSUB_TOPIC}" \
    --region "${GCP_PROJECT_REGION}" \
    --runtime go113 \
    --quiet \
    --memory "512MB"
) &

(
  echo "Deploying FnMessengerWebhook"
  gcloud functions deploy FnMessengerWebhook \
    --trigger-http --allow-unauthenticated \
    --region "${GCP_PROJECT_REGION}" \
    --runtime go113 \
    --memory "512MB"
) &

wait
