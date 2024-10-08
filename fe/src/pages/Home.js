import React, { useState, useEffect } from 'react';

const Home = () => {
  const [healthStatus, setHealthStatus] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    // Fetch the health check from the API
    const fetchHealthStatus = async () => {
      try {
        const response = await fetch('http://localhost:8080/health'); // Adjust the URL if needed
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.text(); // Assuming the API returns plain text
        setHealthStatus(data);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchHealthStatus();
  }, []);

  return (
    <div>
      <h1>Home Page</h1>
      <p>Welcome to the home page!</p>

      {/* Loading spinner or message */}
      {loading && <p>Loading health status...</p>}

      {/* Display error message */}
      {error && <p style={{ color: 'red' }}>Error: {error}</p>}

      {/* Display health status once loaded */}
      {!loading && !error && (
        <div>
          <h2>Health Status</h2>
          <p>{healthStatus}</p>
        </div>
      )}
    </div>
  );
};

export default Home;
