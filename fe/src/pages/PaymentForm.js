import React, { useState } from 'react';
import axios from 'axios';

const PaymentForm = () => {
  const [paymentData, setPaymentData] = useState({
    transaction_id: '',
    usage: '',
    description: '',
    amount: '',
    currency: '',
    consumer_id: '',
    customer_email: '',
    customer_phone: '',
  });

  const handleInputChange = (e) => {
    setPaymentData({ ...paymentData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/create-payment', paymentData);
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
        type="text"
        name="usage"
        value={paymentData.usage}
        onChange={handleInputChange}
        placeholder="Usage"
        required
      />
      <input
        type="text"
        name="description"
        value={paymentData.description}
        onChange={handleInputChange}
        placeholder="Description"
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
      <input
        type="text"
        name="currency"
        value={paymentData.currency}
        onChange={handleInputChange}
        placeholder="Currency"
        required
      />
      <input
        type="text"
        name="consumer_id"
        value={paymentData.consumer_id}
        onChange={handleInputChange}
        placeholder="Consumer ID"
        required
      />
      <input
        type="email"
        name="customer_email"
        value={paymentData.customer_email}
        onChange={handleInputChange}
        placeholder="Customer Email"
        required
      />
      <input
        type="tel"
        name="customer_phone"
        value={paymentData.customer_phone}
        onChange={handleInputChange}
        placeholder="Customer Phone"
        required
      />
      <button type="submit">Pay Now</button>
    </form>
  );
};

export default PaymentForm;