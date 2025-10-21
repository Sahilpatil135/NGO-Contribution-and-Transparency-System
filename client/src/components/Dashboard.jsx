import { useAuth } from '../contexts/AuthContext';
import './Dashboard.css';
import { Link } from "react-router-dom";
import heroimg from "../assets/img1.png"

const Dashboard = () => {
  const { user, logout, fetchCurrentUser } = useAuth();

  const handleLogout = () => {
    logout();
  };

  if (!user) {
    return null;
  }

  console.log(user)

  // return (
  //   <div className="dashboard">
  //     <header className="dashboard-header">
  //       <div className="dashboard-header-content">
  //         <h1>NGO Contribution & Transparency System</h1>
  //         <div className="user-info">
  //           <span>Welcome, {user.name}</span>
  //           <button onClick={handleLogout} className="logout-button">
  //             Logout
  //           </button>
  //         </div>
  //       </div>
  //     </header>

  //     <main className="dashboard-main">
  //       <div className="dashboard-content">
  //         <div className="welcome-section">
  //           <h2>Welcome to your Dashboard</h2>
  //           <p>This is where you'll manage your NGO contributions and view transparency reports.</p>
  //         </div>

  //         <div className="features-grid">
  //           <Link to="/makeContribution" style={{ textDecoration: "none" }}>
  //             <div className="feature-card">
  //               <h3>Make Contributions</h3>
  //               <p>Donate to various causes and track your contributions</p>
  //             </div>
  //           </Link>

  //           <div className="feature-card">
  //             <h3>View Reports</h3>
  //             <p>Access detailed transparency reports and impact metrics</p>
  //           </div>

  //           <div className="feature-card">
  //             <h3>Manage Profile</h3>
  //             <p>Update your personal information and preferences</p>
  //           </div>
  //         </div>
  //       </div>
  //     </main>
  //   </div>
  // );

  return (
    <div className="home-container">
      {/* Navbar */}
      <nav className="navbar">
        <div className= "navbar-left">
          <div className="logo">
            <span className="logo-icon">●</span> CharityLight
          </div>
          <ul className="nav-links">
            <li>Home</li>
            <li>About Us</li>
            <li>Campaigns</li>
            <li>Blog</li>
            <li>Contact</li>
          </ul>
        </div>

        <div className="nav-btns">
          <button className="donate-btn">Login</button>
          <button className="donate-btn">Logout</button>
        </div>

      </nav>

      {/* Hero Section */}
      <section className="hero">
        <div className="hero-text">
          <h1>
            Your <span className="highlight">Support</span>
            <br />
            Can Transform
            <br />
            Communities
          </h1>

          <div className="donation-count">
            <h2>5000+</h2>
            <p>People donated</p>
          </div>

          <p className="hero-description">
            Empowering communities, transforming lives. Your support can bring
            hope to those in need.
          </p>

          <div className="hero-buttons">
            <Link to="/makeContribution" style={{ textDecoration: "none" }}>
              <button className="donate-btn">Donate Now</button>
            </Link>            
            <button className="get-involved-btn">Get Involved</button>
          </div>
        </div>

        <div className="hero-images">
          <img
            src={heroimg}
            alt="volunteer1"
            className="hero-img"
          />
          {/* <img
            src="https://images.unsplash.com/photo-1593113598332-cd68f3e3a4c2?auto=format&fit=crop&w=600&q=80"
            alt="volunteer2"
            className="hero-img"
          /> */}
        </div>
      </section>
    </div>
  );

};

export default Dashboard;
