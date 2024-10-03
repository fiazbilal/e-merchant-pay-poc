import React from 'react';
import './Modal.css'; // You can use the same CSS for both modals

const handleBackToPayment = () => {
  window.location.href = 'http://localhost:3000/payment';
};

const FailureModal = () => (
  <div className="modal">
    <div className="modal-content">
      <h2>Payment Failed</h2>
      <p>There was an error processing your payment. Please try again.</p>
      <button onClick={handleBackToPayment}>Close</button>
    </div>
  </div>
);

export default FailureModal;
