import { createContext, useContext, useEffect, useState } from "react";
import {
    getUserProfile,
    refreshTokenIfNeeded,
    logout as apiLogout,
    loginWithEmailPassword,
    registerUser,
    loginWithGoogle,
    handleGoogleCallback,
} from "../services/auth";

interface AuthState {
    userData: any,
    isAuthenticated: boolean,
    isLoading: boolean,
    authError: string,
    registerUser: (name: string, email: string, password: string, confirmPassword: string) => Promise<boolean>,
    loginWithEmail: (email: string, password: string) => Promise<boolean>,
    loginWithGoogle: () => Promise<boolean>,
    processGoogleCallback: (code: any) => Promise<boolean>,
    logout: () => Promise<void>,
}

const AuthContext = createContext<AuthState | null>(null);

export const useAuth = () => useContext(AuthContext);

export const AuthProvider = ({ children }: any) => {
    const [userData, setUser] = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const [authError, setAuthError] = useState("");
    const isAuthenticated = userData != null;

    const setLogin = async (userData: any) => {
        setUser(userData);
        setAuthError("");
    };

    const clearLogin = async (err: string) => {
        setUser(null);
        setAuthError(err);
    };

    useEffect(() => {
        const initAuth = async () => {
            setIsLoading(true);

            try {
                const isTokenValid = await refreshTokenIfNeeded();

                if (isTokenValid) {
                    const userData = await getUserProfile();
                    setLogin(userData);
                } else {
                    clearLogin("Please log in to continue.");
                }
            } catch (error) {
                console.error(error);
                clearLogin("Please log in to continue.");
            } finally {
                setIsLoading(false);
            }
        };

        initAuth();
    }, []);

    useEffect(() => {
        if (isAuthenticated) {
            const refreshTimer = setInterval(
                async () => {
                    try {
                        await refreshTokenIfNeeded();
                    } catch (error) {
                        console.error(error);
                        clearLogin("Token refresh failed.");
                    }
                },
                30 * 60 * 1000,
            );

            return () => clearInterval(refreshTimer);
        }
    }, [isAuthenticated]);

    const logout = async () => {
        try {
            await apiLogout();
        } catch (error) {
            console.error(error);
        } finally {
            clearLogin("");
        }
    };

    const startGoogleLogin = async () => {
        try {
            const url = await loginWithGoogle();
            window.location.href = url;
            return true;
        } catch (error) {
            console.error(error);
            throw error;
        }
    };

    const processGoogleCallback = async (token: any) => {
        try {
            const userData = await handleGoogleCallback(token);
            setLogin(userData);
            return true;
        } catch (error) {
            console.error(error);
            throw error;
        }
    };

    const loginWithEmail = async (email: string, password: string) => {
        try {
            const userData = await loginWithEmailPassword(email, password);
            setLogin(userData);
            return true;
        } catch (error) {
            console.error(error);
            throw error;
        }
    };

    const register = async (name: string, email: string, password: string, confirmPassword: string) => {
        try {
            if (password != confirmPassword) {
                throw new Error("Passwords do not match");
            }
            const userData = await registerUser({
                name,
                email,
                password,
                confirm_password: confirmPassword,
            });
            setLogin(userData);
            return true;
        } catch (error) {
            console.error(error);
            throw error;
        }
    };

    let authState: AuthState = {
        userData: userData,
        isAuthenticated: isAuthenticated,
        isLoading: isLoading,
        authError: authError,
        registerUser: register,
        loginWithEmail: loginWithEmail,
        loginWithGoogle: startGoogleLogin,
        processGoogleCallback: processGoogleCallback,
        logout: logout,
    };

    return (
        <AuthContext.Provider value={authState}>
            {children}
        </AuthContext.Provider>
    );
}