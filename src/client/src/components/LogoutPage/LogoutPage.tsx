import { Navigate } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';

export function LogoutPage() {
    const authState = useAuth();
    authState?.logout();
    return (
        <Navigate to="/" />
    );
}
