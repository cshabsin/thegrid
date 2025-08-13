import { initializeApp } from "firebase/app";
import { getAuth, GoogleAuthProvider, onAuthStateChanged, signInWithPopup, signOut } from "firebase/auth";

let app;
let auth;

window.firebaseAuth = {
    initializeApp: (config) => {
        app = initializeApp(config);
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
