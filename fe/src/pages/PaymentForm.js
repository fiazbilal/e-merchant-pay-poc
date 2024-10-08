import React, { useState } from 'react';
import axios from 'axios';

const PaymentForm = () => {
  const [paymentData, setPaymentData] = useState({
    transaction_id: '123', // Transaction ID can still be dynamic
    amount: 120.12,
    currency: 'EUR', // Default currency to USD
    notification_url: 'http://localhost:8080/webhook-notification', 
    return_success_url: 'http://localhost:3000/payment-success', 
    return_failure_url: 'http://localhost:3000/payment-failure', 
    transaction_types: [{ name: 'sale' }],
  });

  const handleInputChange = (e) => {
    setPaymentData({ ...paymentData, [e.target.name]: e.target.value });
  };

  const handleCurrencyChange = (e) => {
    setPaymentData({ ...paymentData, currency: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/create-payment', paymentData);
      console.log("Resp: ",response)
      const { redirect_url } = response.data;
      window.location.href = redirect_url;
    } catch (error) {
      console.error('Error creating payment:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        name="transaction_id"
        value={paymentData.transaction_id}
        onChange={handleInputChange}
        placeholder="Transaction ID"
        required
      />
      <input
        type="number"
        name="amount"
        value={paymentData.amount}
        onChange={handleInputChange}
        placeholder="Amount"
        required
      />
      <select name="currency" value={paymentData.currency} onChange={handleCurrencyChange} required>
        <option value="USD">USD</option>
        <option value="EUR">EUR</option>
      </select>
      <button type="submit">Pay Now</button>
    </form>
  );
};

export default PaymentForm;
