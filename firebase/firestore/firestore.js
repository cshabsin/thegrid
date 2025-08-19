import { getFirestore, doc, setDoc } from "firebase/firestore";

let db;

window.firestore = {
    initialize: (app) => {
        db = getFirestore(app);
    },
    createGame: (game) => {
        return setDoc(doc(db, "games", game.id), game);
    },
};