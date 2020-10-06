import firebase from 'firebase/app'

export function configureOneSignal() {
  firebase.auth().onAuthStateChanged((user) => {
    if (user) {
      setOnesignalExternalUserId(user.uid)
    }
  })
}

function setOnesignalExternalUserId(uid: string) {
  const oneSignalInstance = (window as any).OneSignal
  if (oneSignalInstance) {
    oneSignalInstance.push(() => {
      oneSignalInstance.setExternalUserId(uid);
    });
  }
}
