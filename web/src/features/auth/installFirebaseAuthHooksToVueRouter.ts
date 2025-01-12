import type VueRouter from 'vue-router';
import 'firebase/auth';
import firebase from 'firebase/app';

function getCurrentFirebaseUser() {
  return new Promise((resolve, reject) => {
    const unsubscribe = firebase.auth().onAuthStateChanged((user) => {
      unsubscribe();
      resolve(user);
    }, reject);
  });
}


export function installFirebaseAuthHooksToVueRouter(router: VueRouter) {
  router.beforeEach(async (to, _, next) => {
    const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);
    if (requiresAuth && !await getCurrentFirebaseUser()) {
      next('/login');
    } else {
      next();
    }
  });
}

const router!: VueRouter
const routes = [
  { path: '/members', component: MembersPage, meta: { requiresAuth: true } },
  { path: '/login', component: LoginPage },
  { path: '/logout', component: LogoutPage },
]
router.beforeEach(async (to, _, next) => {
  const isLoggedIn = firebase.auth().currentUser !== null
  const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);
  if (requiresAuth && !await getCurrentFirebaseUser()) {
    next('/login');
  } else {
    next();
  }
});
