import { useAuth } from '../contexts/AuthContext';
import './Dashboard.css';

const Dashboard = () => {
  const { user, logout, fetchCurrentUser } = useAuth();

  const handleLogout = () => {
    logout();
  };

  if (!user) {
    return null;
  }

  console.log(user)

  return (
    <div className="dashboard">
      <header className="dashboard-header">
        <div className="dashboard-header-content">
          <h1>NGO Contribution & Transparency System</h1>
          <div className="user-info">
            <span>Welcome, {user.name}</span>
            <button onClick={handleLogout} className="logout-button">
              Logout
            </button>
          </div>
        </div>
      </header>

      <main className="dashboard-main">
        <div className="dashboard-content">
          <div className="welcome-section">
            <h2>Welcome to your Dashboard</h2>
            <p>This is where you'll manage your NGO contributions and view transparency reports.</p>
          </div>

          <div className="features-grid">
            <div className="feature-card">
              <h3>Make Contributions</h3>
              <p>Donate to various causes and track your contributions</p>
            </div>

            <div className="feature-card">
              <h3>View Reports</h3>
              <p>Access detailed transparency reports and impact metrics</p>
            </div>

            <div className="feature-card">
              <h3>Manage Profile</h3>
              <p>Update your personal information and preferences</p>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
};

export default Dashboard;
