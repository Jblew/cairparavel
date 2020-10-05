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
      return request.time < resource.data.votingTime
    }

    function inSignupPeriod() {
      return request.time > resource.data.votingTime && request.time < resource.data.signupTime
    }

    match /events/{eventId} {
      allow read: if isAuthenticatedMember();
      allow create: if isAuthenticatedMember() && ensureOwnerUid() && ensureEventCreateFields();
      allow delete: if false;
      allow update: if ensureEventCreateFields() || resource.data.ownerUid == request.auth.uid;

      match /votes/{uid} {
        allow read: if isAuthenticatedMember();
        allow delete: if uid == request.auth.uid && inVotingModificationPeriod();
        allow create, update: if uid == request.auth.uid && inVotingModificationPeriod();
      }

      match /signedMembers/{uid} {
        allow read: if isAuthenticatedMember();
        allow delete: if uid == request.auth.uid && inSignupPeriod();
        allow create, update: if uid == request.auth.uid && inSignupPeriod();
      }
    }
  `
}

const data = getRules()
fs.writeFileSync(__dirname + '/firestore.rules', data)
console.log('Regenerated firestore rules')
