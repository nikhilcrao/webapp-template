import { createContext, useContext, useEffect, useState } from "react";
import {
    registerUser,
    loginWithEmail,
    refreshTokenIfNeeded,
    handleGoogleCallback,
    logout as apiLogout,
    getUserProfile,
} from "../services/auth";

const AuthContext = createContext({
    userData: null,
    isAuthenticated: false,
    isLoading: false,
    handleRegister: (_n: string, _e: string, _p: string, _c: string) => {
        return Promise.resolve(false);
    },
    handleLogin: (_e: string, _p: string) => {
        return Promise.resolve(false);
    },
    handleLogout: () => { },
    processGoogleCallback: (_code: string) => {
        return Promise.resolve(false);
    },
});

export const useAuth = () => useContext(AuthContext);

export const AuthProvider = ({ children }: any) => {
    const [userData, setUser] = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const isAuthenticated = userData != null;

    const setLogin = async (userData: any) => {
        setUser(userData);
    };

    const clearLogin = async () => {
        setUser(null);
    };

    // Initialize the auth state.
    useEffect(() => {
        const initAuth = async () => {
            setIsLoading(true);

            try {
                const validToken = await refreshTokenIfNeeded();
                if (validToken) {
                    const userData = await getUserProfile();
                    setLogin(userData);
                } else {
                    clearLogin();
                }
            } catch (error) {
                console.error(error);
                clearLogin();
            } finally {
                setIsLoading(false);
            }
        };

        initAuth();
    }, []);

    // Refresh the token every 30 mins in the background.
    useEffect(() => {
        if (isAuthenticated) {
            const refreshTimer = setInterval(
                async () => {
                    try {
                        await refreshTokenIfNeeded();
                    } catch (error) {
                        console.error(error);
                        clearLogin();
                    }
                },
                30 * 60 * 1000,  // 30 mins
            );

            return () => clearInterval(refreshTimer);
        }
    }, [isAuthenticated]);

    const handleLogin = async (email: string, password: string) => {
        try {
            const userData = await loginWithEmail(email, password);
            setLogin(userData);
            return true;
        } catch (error) {
            console.error(error);
            throw error;
        }
    };

    const handleRegister = async (name: string, email: string, password: string, confirmPassword: string) => {
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

    const handleLogout = async () => {
        try {
            await apiLogout();
        } catch (error) {
            console.error(error);
        } finally {
            clearLogin();
        }
    };


    const processGoogleCallback = async (code: any) => {
        try {
            const userData = await handleGoogleCallback(code);
            setLogin(userData);
            return true;
        } catch (error) {
            console.error(error);
            throw error;
        }
    };

    return (
        <AuthContext.Provider value={{
            userData,
            isAuthenticated,
            isLoading,
            handleRegister,
            handleLogin,
            handleLogout,
            processGoogleCallback,
        }}>
            {children}
        </AuthContext.Provider >
    );
}