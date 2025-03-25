import axios from "axios";

const api = axios.create({
    baseURL: "/api",
    headers: {
        "Content-Type": "application/json",
    },
})

api.inteceptors.request.use(
    (config) => {
        const token = localStorage.getItem("authToken");
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    },
);

api.inteceptors.response.use(
    (response) => {
        return response;
    },
    (error) => {
        const { response } = error;

        if (response && response.status === 401) {
            if (!window.location.pathname.includes("/login")) {
                localStorage.removeItem("authToken");
                window.location.href = "/login";
            }
        }

        const errorMessage =
            (response && response.data && response.data.errorMessage) ||
            "Something went wrong. Please try again later.";
        
            return Promise.reject({
                ...error,
                userMessage: errorMessage,
            });
    },
);

export default api;