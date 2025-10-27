import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { AuthProvider, useAuth } from './contexts/AuthContext';
import Login from './components/auth/Login';
import Signup from './components/auth/Signup';
import OAuthCallback from './pages/OAuthCallback';
import './App.css';
import Navbar from './components/Navbar';
import Footer from './components/Footer';
import HomePage from './Pages/HomePage';
import ContributionsPage from './Pages/ContributionsPage';
import CampaignPage from './Pages/CampaignPage';
import CheckoutPage from './Pages/CheckoutPage';
import DonationTypePage from './Pages/DonationTypePage';
import NgoRegistration from './Pages/NgoRegistration';
import CreateCampaign from './Pages/CreateCampaign';

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
        <Route path="/makeContribution" element={isAuthenticated ? <ContributionsPage /> : <Navigate to="/login" replace />} />
        <Route path="/campaign/:causeID" element={isAuthenticated ? <CampaignPage /> : <Navigate to="/login" replace />} />
        <Route path="/checkout" element={isAuthenticated ? <CheckoutPage /> : <Navigate to="/login" replace />} />
        <Route path="/makeContribution/:category/:slug" element={isAuthenticated ? <DonationTypePage /> : <Navigate to="/login" replace />} />
        <Route path="/ngoRegistration" element={isAuthenticated ? <NgoRegistration /> : <Navigate to="/login" replace />} />
        <Route path="/createCampaign" element={isAuthenticated ? <CreateCampaign /> : <Navigate to="/login" replace />} />
      </Routes>

      {isAuthenticated && <Footer />}
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
