declare module 'auth';

export function registerUser(any): Promise<any>;
export function loginWithEmail(string, string): Promise<any>;
export function handleGoogleCallback(any): any;
export function refreshTokenIfNeeded(): Promise<boolean>;
export function logout(): any;
export function getUserProfile(): any;