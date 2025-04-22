import '@mantine/core/styles.css';
import '@mantine/dates/styles.css';
import '@mantine/charts/styles.css';
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from 'react-router-dom';
import { Center, Loader, MantineProvider } from '@mantine/core';
import { AppLayout } from './components/layout/AppLayout/AppLayout';
import { AuthProvider } from './contexts/AuthContext';
import { LoginPage } from './pages/Login/Login';
import { useAuth } from './contexts/AuthContext';
import { LogoutPage } from './pages/Logout/Logout';
import { GoogleOAuthProvider } from '@react-oauth/google';
import { Dashboard } from './pages/Dashboard/Dashboard';
import { PageNotFound } from './pages/PageNotFound/PageNotFound';

const ProtectedRoute = ({ children }: { children: any }) => {
  const { isLoading, isAuthenticated } = useAuth();
  if (isLoading) {
    return (
      <Center>
        <Loader />
      </Center>
    );
  }
  if (isAuthenticated) {
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
        <GoogleOAuthProvider clientId="848998314068-eprp0d0lfln4vmpo1iktnvids4srahvq.apps.googleusercontent.com">
          <Router>
            <Routes>
              {/* Auth Routes */}
              <Route path="/login" element={<LoginPage />} />
              <Route path="/logout" element={
                <ProtectedRoute>
                  <LogoutPage />
                </ProtectedRoute>
              } />

              {/* App Layout Wrapper */}
              <Route element={<AppLayout />}>
                <Route path="/" element={
                  <Dashboard />
                } />
              </Route>

              {/* Fallback */}
              <Route path="*" element={<PageNotFound />} />
            </Routes>
          </Router>
        </GoogleOAuthProvider>
      </AuthProvider>
    </MantineProvider >
  );
}