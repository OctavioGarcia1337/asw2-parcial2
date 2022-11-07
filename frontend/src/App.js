import React, { useState } from "react";
import { BrowserRouter as Router, Route, Routes, Link} from "react-router-dom";
import "./css/App.css";
import Home from "./Home"
import Product from "./Product"

function App(){
return (
    <Router>
      <Routes>
        <Route exact path = "/" element={<Home/>}/>
        <Route path= "/product" element={<Product/>}/>
      </Routes>
    </Router>
  );
}

export default App;
