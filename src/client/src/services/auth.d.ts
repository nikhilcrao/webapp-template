declare module 'auth';

export function loginWithEmailPassword(string, string): Promise<any>;
export function refreshTokenIfNeeded(): Promise<boolean>;
export function registerUser(any): Promise<any>;
export function loginWithGoogle(): any;
export function handleGoogleCallback(any): any;
export function getUserProfile(): any;
export function logout(): any;