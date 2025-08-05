window.firebaseAuth = {
    initializeApp: (config) => {
        firebase.initializeApp(config);
    },
    signIn: () => {
        const provider = new firebase.auth.GoogleAuthProvider();
        firebase.auth().signInWithPopup(provider);
    },
    signOut: () => {
        firebase.auth().signOut();
    },
    onAuthStateChanged: (callback) => {
        firebase.auth().onAuthStateChanged(callback);
    },
};