import { initializeApp } from "firebase/app";

window.firebase = {
    initializeApp: (config) => {
        return initializeApp(config);
    },
};
