import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import App from "./App"
import RealtimePage from "./pages/Realtime"

const MyRoutes = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<RealtimePage />} />
      </Routes>
    </BrowserRouter>
  );
};

export default MyRoutes;
