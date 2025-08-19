import { getAuth, GoogleAuthProvider, onAuthStateChanged, signInWithPopup, signOut } from "firebase/auth";

let auth;

window.firebaseAuth = {
    initialize: (app) => {
        auth = getAuth(app);
    },
    signIn: () => {
        const provider = new GoogleAuthProvider();
        signInWithPopup(auth, provider);
    },
    signOut: () => {
        signOut(auth);
    },
    onAuthStateChanged: (callback) => {
        onAuthStateChanged(auth, callback);
    },
};