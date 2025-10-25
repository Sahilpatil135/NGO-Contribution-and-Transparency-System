import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { AuthProvider, useAuth } from './contexts/AuthContext';
import Login from './components/auth/Login';
import Signup from './components/auth/Signup';
import OAuthCallback from './pages/OAuthCallback';
import './App.css';
import Navbar from './components/Navbar';
import HomePage from './Pages/HomePage';
import ContributionsPage from './Pages/ContributionsPage';
import CampaignPage from './Pages/CampaignPage';
import CheckoutPage from './Pages/CheckoutPage';
import DonationTypePage from './Pages/DonationTypePage';

const AppRoutes = () => {
  const { isAuthenticated, isLoading } = useAuth();

  if (isLoading) {
    return (
      <div className="loading-container">
        <div className="loading-spinner"></div>
        <p>Loading...</p>
      </div>
    );
  }

  return (
    <>
      {isAuthenticated && <Navbar />}

      <Routes>
        <Route
          path="/login"
          element={
            isAuthenticated ? <Navigate to="/" replace /> : <Login />
          }
        />
        <Route
          path="/signup"
          element={
            isAuthenticated ? <Navigate to="/" replace /> : <Signup />
          }
        />
        <Route
          path="/"
          element={
            isAuthenticated ? <HomePage /> : <Navigate to="/login" replace />
          }
        />
        <Route
          path="/auth/callback"
          element={<OAuthCallback />}
        />
        {/* <Route
        path="/"
        element={
          <Navigate to={isAuthenticated ? "/" : "/login"} replace />
        }
      /> */}
        <Route path="/makeContribution" element={isAuthenticated ? <ContributionsPage /> : <Navigate to="/login" replace />} />
        <Route path="/campaign/:id" element={isAuthenticated ? <CampaignPage /> : <Navigate to="/login" replace />} />
        <Route path="/checkout" element={isAuthenticated ? <CheckoutPage /> : <Navigate to="/login" replace />} />
        <Route path="/makeContribution/:id" element={isAuthenticated ? <DonationTypePage /> : <Navigate to="/login" replace />} />
      </Routes>
    </>
  );
};

function App() {
  return (
    <AuthProvider>
      <Router>
        <div className="app">
          {/* <Navbar /> */}
          <AppRoutes />
        </div>
      </Router>
    </AuthProvider>
  );
}

export default App;
