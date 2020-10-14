const fs = require('fs')

function getRules() {
  return `
rules_version = '2';

service cloud.firestore {
  match /databases/{database}/documents {
    match /{document=**} {
      allow write: if false;
      allow read: if false;
    }

    ///////////////////////
    // Utility functions //
    ///////////////////////

    function isAuthenticated() {
        return request.auth != null;
    }

    function notUpdating(field) {
      return !(field in request.resource.data)
        || resource.data[field] == request.resource.data[field]
    }

    ${rulesForEnv('test')}
    ${rulesForEnv('prod')}
  }
}
  `
}

function rulesForEnv(envName) {
  const basePath = `/databases/$(database)/documents/envs/${envName}`
  return `
  match /envs/${envName} {
    ${roles({ basePath })}


    ${users()}


    ${events()}


    ${comments()}


    ${notifications()}
  }
  `
}

function roles(opts) {
  if (!opts.basePath) throw new Error('Missing opts.basePath')
  return `
    /////////////////
    // Role system //
    /////////////////

    function userHasRole(role, uid) {
        return isAuthenticated()
            && exists(${opts.basePath}/roles/$(role)/uids/$(uid));
    }

    match /roles/{roleName}/uids/{uid} {
      allow write: if false;
      allow read: if isAuthenticated()
    }

    // Role: leader
    match /roles/leader/uids/{uid} {
      allow write: if false;
    }

    // Role: member
    match /roles/member/uids/{uid} {
      allow write: if userHasRole("leader", request.auth.uid);
    }
  `
}

function users() {
  return `
    ///////////
    // Users //
    ///////////

    match /users/{uid} {
      allow write: if false;
      allow read: if userHasRole("leader", request.auth.uid);
    }

    function isAuthenticatedMember() {
      return userHasRole("member", request.auth.uid);
    }
  `
}

function events() {
  return `
    ////////////
    // Events //
    ////////////

    function ensureOwnerUid() {
      return isAuthenticatedMember() && request.resource.data.ownerUid == request.auth.uid
    }

    function ensureEventCreateFields() {
      return notUpdating('votes') && notUpdating('signedMembers')
    }

    function inVotingModificationPeriod() {
      //
      return request.time.toMillis() < resource.data.votingTime
    }

    function inSignupPeriod() {
      return request.time.toMillis() > 0 // resource.data.votingTime // && request.time.toMillis() < resource.data.signupTime
    }

    match /events/{eventId} {
      allow read: if isAuthenticatedMember();
      allow create: if isAuthenticatedMember() && ensureOwnerUid() && ensureEventCreateFields();
      allow delete: if resource.data.ownerUid == request.auth.uid;
      allow update: if ensureEventCreateFields() || resource.data.ownerUid == request.auth.uid;

      match /votes/{uid} {
        allow read: if isAuthenticatedMember();
        allow delete: if uid == request.auth.uid; // && inVotingModificationPeriod();
        allow create, update: if uid == request.auth.uid; // && inVotingModificationPeriod();
      }

      match /signedMembers/{uid} {
        allow read: if isAuthenticatedMember();
        allow delete: if uid == request.auth.uid && inSignupPeriod();
        allow create, update: if uid == request.auth.uid && inSignupPeriod();
      }
    }

    match /event_observers/{eventId}/uids/{uid} {
      allow read: if uid == request.auth.uid;
      allow write: if uid == request.auth.uid;
    }

    match /event_laststate/{eventId} {
      allow read: if false;
      allow write: if false;
    }
  `
}

function comments() {
  return `
    //////////////
    // Comments //
    //////////////

    match /event_comments/{eventId}/messages/{commentId} {
      allow create: if request.resource.data.authorUid == request.auth.uid;
      allow delete: if resource.data.authorUid == request.auth.uid;
      allow update: if false;
      allow read: if isAuthenticatedMember();
    }
  `
}

function notifications() {
  return `
    //////////////////
    // Notifications //
    //////////////////

    match /notifications/{uid}/messenger_queue/{notificationId} {
      allow write: if false;
      allow read: if isAuthenticatedMember() && uid == request.auth.uid;
    }

    match /notifications/{uid}/history/{notificationId} {
      allow write: if false;
      allow read: if isAuthenticatedMember() && uid == request.auth.uid;
    }

    match /notifications_allow/{type}/uids/{uid} {
      allow create: if uid == request.auth.uid;
      allow delete: if uid == request.auth.uid;
    }

    match /messenger_recipients/{uid} {
      allow create: if false;
      allow delete: if false;
    }
  `
}

const data = getRules()
fs.writeFileSync(__dirname + '/firestore.rules', data)
console.log('Regenerated firestore rules')
