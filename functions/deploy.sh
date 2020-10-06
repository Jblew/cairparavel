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

gcloud functions deploy FnOnUserCreated \
  --trigger-event providers/firebase.auth/eventTypes/user.create \
  --trigger-resource "${GCP_PROJECT_ID}" \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --memory "1024MB"

gcloud functions deploy FnOnNotificationCreatedPushNotification \
  --trigger-event providers/cloud.firestore/eventTypes/document.create \
  --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/notifications/{uid}/push/{notificationId}" \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --memory "512MB"

gcloud functions deploy FnOnNotificationsEnabledSendWelcome \
  --trigger-event providers/cloud.firestore/eventTypes/document.create \
  --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/notification_playerids/{uid}" \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --memory "512MB"

gcloud functions deploy FnOnCommentAddedNotifyObservers \
  --trigger-event providers/cloud.firestore/eventTypes/document.create \
  --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/comments/{resId}/messages/{commentId}" \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --memory "512MB"

gcloud functions deploy FnOnEventCreatedNotify \
  --trigger-event providers/cloud.firestore/eventTypes/document.create \
  --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/events/{eventId}" \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --memory "512MB"

gcloud functions deploy FnOnEventModifiedNotifyObservers \
  --trigger-event providers/cloud.firestore/eventTypes/document.update \
  --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/events/{eventId}" \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --memory "512MB"

gcloud functions deploy FnOnEventMemberSignedUpNotifyObservers \
  --trigger-event providers/cloud.firestore/eventTypes/document.create \
  --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/events/{eventId}/signedMembers/{uid}" \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --memory "512MB"

gcloud functions deploy FnOnEventMemberSignedOutNotifyObservers \
  --trigger-event providers/cloud.firestore/eventTypes/document.delete \
  --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/events/{eventId}/signedMembers/{uid}" \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --memory "512MB"

gcloud functions deploy FnOnEventVotedNotifyObservers \
  --trigger-event providers/cloud.firestore/eventTypes/document.create \
  --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/events/{eventId}/votes/{uid}" \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --memory "512MB"

gcloud functions deploy FnOnEventVoteDeletedNotifyObservers \
  --trigger-event providers/cloud.firestore/eventTypes/document.delete \
  --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/events/{eventId}/votes/{uid}" \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --memory "512MB"

gcloud functions deploy FnOnEventVoteDeletedNotifyObservers \
  --trigger-event providers/cloud.firestore/eventTypes/document.delete \
  --trigger-resource "projects/${GCP_PROJECT_ID}/databases/(default)/documents/envs/{env}/events/{eventId}/votes/{uid}" \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --memory "512MB"


CHECK_EVENTS_PUBSUB_TOPIC="cron-check-events-state"
CHECK_EVENTS_PUBSUB_SUBSCRIPTION="sub_check-events-state"
CHECK_EVENTS_SCHEDULER_JOB="job_check-events-state"
gcloud pubsub topics create "${CHECK_EVENTS_PUBSUB_TOPIC}"
gcloud pubsub subscriptions create "${CHECK_EVENTS_PUBSUB_SUBSCRIPTION}" --topic "${CHECK_EVENTS_PUBSUB_TOPIC}"
gcloud alpha scheduler jobs delete "${CHECK_EVENTS_SCHEDULER_JOB}" || echo "Failed to delete scheduler job"
gcloud alpha scheduler jobs create pubsub "${CHECK_EVENTS_SCHEDULER_JOB}" \
  --topic "${CHECK_EVENTS_PUBSUB_TOPIC}" \
  --schedule "*/10 * * * *" \
  --message-body "SCHEDULE"
gcloud functions deploy FnOnCronDispatchEventStateNotifications \
  --trigger-topic "${CHECK_EVENTS_PUBSUB_TOPIC}" \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --memory "512MB"


gcloud functions deploy FnMessengerWebhook \
  --trigger-http --allow-unauthenticated \
  --region "${GCP_PROJECT_REGION}" \
  --runtime go113 \
  --memory "1024MB"
