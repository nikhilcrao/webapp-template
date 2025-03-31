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
import { LogoutPage } from './components/LogoutPage/LogoutPage';
import { GoogleOAuthProvider } from '@react-oauth/google';

const ProtectedRoute = ({ children }: { children: any }) => {
  const authState = useAuth();

  console.log("protected");

  if (authState?.isLoading) {
    console.log("loading");
    return (
      <Loader />
    );
  }

  console.log("authState", authState);

  if (authState?.isAuthenticated) {
    console.log("authenticated");
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
                  <ProtectedRoute>
                    <DashboardPage />
                  </ProtectedRoute>
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