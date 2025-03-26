import '@mantine/core/styles.css';
import '@mantine/dates/styles.css';
import '@mantine/charts/styles.css';

import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from 'react-router-dom';
import { Loader, MantineProvider } from '@mantine/core';
import { AuthProvider } from './contexts/AuthContext';
import { PageNotFound } from './components/PageNotFound/PageNotFound';
import { LoginPage } from './components/LoginPage/LoginPage';
import { AppLayout } from './components/Layout/AppLayout';
import { useAuth } from './contexts/AuthContext';
import { DashboardPage } from './components/DashboardPage/DashboardPage';
import { ProfilePage } from './components/ProfilePage/ProfilePage';
import { LogoutPage } from './components/LogoutPage/LogoutPage';

const ProtectedRoute = ({ children }: { children: any }) => {
  const authState = useAuth();

  if (authState?.isLoading) {
    return (
      <Loader />
    );
  }

  if (authState?.isAuthenticated) {
    return children;
  }

  return (
    <Navigate to="/login" />
  );
}

export default function App() {
  return (
    <MantineProvider defaultColorScheme="auto">
      <AuthProvider>
        <Router>
          <Routes>
            {/* Auth Routes */}
            <Route path="/login" element={<LoginPage />} />
            <Route path="/logout" element={<LogoutPage />} />

            {/* App Layout Wrapper */}
            <Route element={<AppLayout />}>
              <Route path="/" element={
                <ProtectedRoute>
                  <DashboardPage />
                </ProtectedRoute>} />

              <Route path="/profile" element={
                <ProtectedRoute>
                  <ProfilePage />
                </ProtectedRoute>} />
            </Route>

            {/* Fallback */}
            <Route path="*" element={<PageNotFound />} />
          </Routes>
        </Router>
      </AuthProvider>
    </MantineProvider >
  );
}