// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
// Cypress.Commands.add("login", (email, password) => { ... })
//
//
// -- This is a child command --
// Cypress.Commands.add("drag", { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add("dismiss", { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This is will overwrite an existing command --
// Cypress.Commands.overwrite("visit", (originalFn, url, options) => { ... })

import firebase from "firebase/app";
import "firebase/auth";
import "firebase/database";
import "firebase/firestore";
import { attachCustomCommands } from "cypress-firebase";

const FIREBASE_CONFIG = {
  apiKey: "AIzaSyAG3Q-B7U5rBgu5O5KJ5XeGlyHm-8m8Cis",
  authDomain: "cairparavelapp.firebaseapp.com",
  databaseURL: "https://cairparavelapp.firebaseio.com",
  projectId: "cairparavelapp",
  storageBucket: "cairparavelapp.appspot.com",
  messagingSenderId: "486328450948",
  appId: "1:486328450948:web:e78910b40a9d967f3e27d1"
};

firebase.initializeApp(FIREBASE_CONFIG);

attachCustomCommands({ Cypress, cy, firebase });
