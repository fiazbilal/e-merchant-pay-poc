import React from 'react';
import './Modal.css'; // You can use the same CSS for both modals

const handleBackToPayment = () => {
  window.location.href = 'http://localhost:3000/payment';
};

const SuccessModal = () => (
  <div className="modal">
    <div className="modal-content">
      <h2>Payment Successful!</h2>
      <p>Your payment has been processed successfully.</p>
      <button onClick={handleBackToPayment}>Close</button>
    </div>
  </div>
);

export default SuccessModal;
