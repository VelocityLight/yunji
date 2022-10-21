import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import RealtimePage from "./pages/Realtime"
import TeamPage from "./pages/Team"

const MyRoutes = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<RealtimePage />} />
        <Route path="/realtime" element={<RealtimePage />} />
        <Route path="/team" element={<TeamPage />} />
      </Routes>
    </BrowserRouter>
  );
};

export default MyRoutes;
