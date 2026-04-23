import { BrowserRouter as Router, Routes, Route, Navigate, useLocation } from 'react-router-dom';
import { AuthProvider, useAuth } from './contexts/AuthContext';
import Login from './components/auth/Login';
import Signup from './components/auth/Signup';
import OAuthCallback from './Pages/OAuthCallback';
import './App.css';
import Navbar from './components/Navbar';
import Footer from './components/Footer';
import HomePage from './Pages/HomePage';
import ContributionsPage from './Pages/ContributionsPage';
import CampaignPage from './Pages/CampaignPage';
import CheckoutPage from './Pages/CheckoutPage';
import DonationTypePage from './Pages/DonationTypePage';
import DonationSuccess from './Pages/DonationSuccess';
import ProfilePage from './Pages/ProfilePage';
import OrganizationAccountsPage from './Pages/Organization/OrganizationAccountsPage';
import NgoRegistration from './Pages/NgoRegistration';
import CreateCampaign from './Pages/CreateCampaign';
import UploadProof from './Pages/Ngo/UploadProof';
import UploadUpdate from './Pages/Ngo/UploadUpdate';
import MobileProofCapture from './Pages/MobileProofCapture';
import BloodDonationPage from './Pages/BloodDonationPage';
import VolunteerPage from './Pages/VolunteerPage';
import AdminDashboard from './Pages/AdminDashboard';
import {
  NgoVerificationManagementPage,
  DisputeResolutionPanelPage,
  CauseActivityMonitoringPage,
  TrustScoreOverviewPage,
  UserNgoManagementPage
} from './Pages/AdminGovernancePages';

const AppRoutes = () => {
  const { isAuthenticated, isLoading } = useAuth();
  const location = useLocation();

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
        <Route path="/campaign/:causeID" element={isAuthenticated ? <CampaignPage key={location.pathname} /> : <Navigate to="/login" replace />} />
        <Route path="/checkout" element={isAuthenticated ? <CheckoutPage /> : <Navigate to="/login" replace />} />
        <Route path="/bloodDonation" element={isAuthenticated ? <BloodDonationPage /> : <Navigate to="/login" replace />} />
        <Route path="/volunteer" element={isAuthenticated ? <VolunteerPage /> : <Navigate to="/login" replace />} />
        <Route path="/donation/success" element={isAuthenticated ? <DonationSuccess /> : <Navigate to="/login" replace />} />
        <Route path="/profile" element={isAuthenticated ? <ProfilePage /> : <Navigate to="/login" replace />} />
        <Route
          path="/organization/accounts"
          element={isAuthenticated ? <OrganizationAccountsPage /> : <Navigate to="/login" replace />}
        />
        <Route
          path="/organization/:organizationId/accounts"
          element={<OrganizationAccountsPage />}
        />
        <Route path="/makeContribution/:category/:slug" element={isAuthenticated ? <DonationTypePage /> : <Navigate to="/login" replace />} />
        {/* <Route path="/admin" element={isAuthenticated ? <AdminDashboard /> : <Navigate to="/login" replace />} /> */}
        <Route path="/admin" element={<AdminDashboard />} />

        <Route path="/admin/ngo-verifications" element={isAuthenticated ? <NgoVerificationManagementPage /> : <Navigate to="/login" replace />} />
        <Route path="/admin/disputes" element={isAuthenticated ? <DisputeResolutionPanelPage /> : <Navigate to="/login" replace />} />
        <Route path="/admin/causes-monitoring" element={isAuthenticated ? <CauseActivityMonitoringPage /> : <Navigate to="/login" replace />} />
        <Route path="/admin/trust-scores" element={isAuthenticated ? <TrustScoreOverviewPage /> : <Navigate to="/login" replace />} />
        <Route path="/admin/user-ngo-management" element={isAuthenticated ? <UserNgoManagementPage /> : <Navigate to="/login" replace />} />

        <Route path="/ngoRegistration" element={<NgoRegistration />} />

        {/* This routes are only for NGOs. */}
        <Route path="/createCampaign" element={isAuthenticated ? <CreateCampaign /> : <Navigate to="/login" replace />} />
        <Route path="/campaign/:causeID/update" element={isAuthenticated ? <UploadUpdate /> : <Navigate to="/login" replace />} />
        {/* <Route path="/uploadProof" element={isAuthenticated ? <UploadProof /> : <Navigate to="/login" replace />} /> */}
        <Route path="/uploadProof/:causeID" element={<UploadProof />} />
        <Route path="/mobile/proof/:sessionID" element={<MobileProofCapture />} />
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
