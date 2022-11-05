import React from "react";
import { Routes, Route, useLocation } from "react-router-dom";
import Header from "./components/Header";
import Home from "./pages/home/Home";
import SignUp from "./pages/sign-up/SignUp";
import Login from "./pages/login/Login";

import {useSelector} from "react-redux";

const App = () => {

  const token = useSelector((state) => state.Auth.token);
  console.log(token);

  return (
    <div className="background">
      <React.Fragment>
        <Header />
        <Routes>
          <Route path="/" element={<Home />} />
          {/* <Route path="/NavTest1" element={<NavTest />} /> */}
          <Route path="/sign-up" element={<SignUp />} />
          <Route path="/login" element={<Login />} />
        </Routes>
      </React.Fragment>
    </div>
  )
}
export default App;