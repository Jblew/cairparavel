import * as firebase from 'firebase/app';
import 'firebase/firestore';
import 'firebase/auth';
import 'firebase/database';

const FIREBASE_CONFIG = {
  apiKey: "AIzaSyAG3Q-B7U5rBgu5O5KJ5XeGlyHm-8m8Cis",
  authDomain: "cairparavelapp.firebaseapp.com",
  databaseURL: "https://cairparavelapp.firebaseio.com",
  projectId: "cairparavelapp",
  storageBucket: "cairparavelapp.appspot.com",
  messagingSenderId: "486328450948",
  appId: "1:486328450948:web:e78910b40a9d967f3e27d1"
};

export function constructFirebaseApp() {
  return firebase.initializeApp(FIREBASE_CONFIG);
}
