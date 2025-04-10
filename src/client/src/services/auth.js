import api from './api';

export async function loginWithEmail(email, password) {
    try {
        const response = await api.post("/auth/login", { email, password });
        if (response.data.token) {
            localStorage.setItem("token", response.data.token);
        }
        return response.data;
    } catch (error) {
        console.error(error);
        throw error;
    }
}

export async function registerUser(userData) {
    try {
        const response = await api.post("/auth/register", userData);
        if (response.data.token) {
            localStorage.setItem("token", response.data.token);
        }
        return response.data;
    } catch (error) {
        console.error(error);
        throw error;
    }
}

export async function handleGoogleCallback(code) {
    try {
        const response = await api.get(`/auth/google/callback?code=${code}`);

        if (response.data && response.data.token) {
            localStorage.setItem("token", response.data.token);
        }

        return response.data;
    } catch (error) {
        console.error(error);
        throw error;
    }
}

export async function refreshTokenIfNeeded() {
    // TODO: Check token expiration and refresh if needed.
    const token = localStorage.getItem("token");
    return Promise.resolve(!!token);
}

export async function logout() {
    try {
        // TODO: this function needs to call a logout endpoint on the server to invalidate the token on the server.
        localStorage.removeItem("token");
        return Promise.resolve();
    } catch (error) {
        console.error(error);
        throw error;
    }
}

// TODO: Move this into a different service (profile.js)
export async function getUserProfile() {
    try {
        return await api.get("/profile");
    } catch (error) {
        console.error(error);
        throw error;
    }
}
