import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import RealtimePage from "./pages/Realtime"
import TeamPage from "./pages/Team"
import BillingTrend from "./pages/BillingTrend"

const MyRoutes = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<RealtimePage />} />
        <Route path="/home/realtime" element={<RealtimePage />} />
        <Route path="/home/team" element={<TeamPage />} />
        <Route path="/home/trend" element={<BillingTrend />} />
      </Routes>
    </BrowserRouter>
  );
};

export default MyRoutes;
