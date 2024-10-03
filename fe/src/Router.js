// src/Router.js
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'; // Updated import
import Header from './components/Header';
import Home from './pages/Home';
import About from './pages/About';
import PaymentForm from './pages/PaymentForm';
import SuccessModal from './pages/SuccessModal';
import FailureModal from './pages/FailureModal';

const AppRouter = () => {
  return (
    <Router>
      <Header />
      <Routes> {/* Use Routes instead of Switch */}
        <Route path="/" element={<Home />} /> {/* Use element prop instead of component */}
        <Route path="/about" element={<About />} />
        <Route path="/payment" element={<PaymentForm/>}/>
        <Route path="/payment-success" element={<SuccessModal/>}/>
        <Route path="/payment-failure" element={<FailureModal/>}/>

        {/* Add more routes as needed */}
      </Routes>
    </Router>
  );
};

export default AppRouter;
