import '@mantine/core/styles.css';
import '@mantine/dates/styles.css';
import '@mantine/charts/styles.css';

import {
  BrowserRouter as Router,
  Routes,
  Route,
} from 'react-router-dom';
import { MantineProvider } from '@mantine/core';
//import { AppLayout } from "./components/Layout/AppLayout";
import { AuthProvider } from './contexts/AuthContext';
{/* import { LoginPage } from './components/Auth/LoginPage'; */}
import { AppLayout } from './components/Layout/AppLayout';
import { PageNotFound } from './components/PageNotFound/PageNotFound';

export default function App() {
  return (
    <MantineProvider defaultColorScheme="auto">
      <AuthProvider>
        <Router>
          <Routes>
            <Route path="/login" element={ <AppLayout /> } />
            <Route path="*" element={ <PageNotFound /> } />
          </Routes>
        </Router>
        {/* <AppLayout /> */}
      </AuthProvider>
    </MantineProvider>
  );
}