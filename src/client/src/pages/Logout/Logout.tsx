import { Navigate } from 'react-router-dom';
import { useAuth } from '../../contexts/AuthContext';

export function LogoutPage() {
    const { handleLogout } = useAuth();
    handleLogout();
    return (
        <Navigate to="/" />
    );
}
