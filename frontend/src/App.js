import React, { useState } from "react";
import { BrowserRouter as Router, Route, Routes, Link} from "react-router-dom";
import "./css/App.css";
import Home from "./Home"
import Login from "./Login"
import Register from "./Register"
import Item from "./Item"
import Sistema from "./Sistema"
import Publication from "./Publication"
import MyComments from "./MyComments"

function App(){
return (
    <Router>
      <Routes>
        <Route exact path = "/" element={<Home/>}/>
        <Route path= "/login" element={<Login/>}/>
        <Route path= "/register" element={<Register/>}/>
        <Route path= "/item" element={<Item/>}/>
        <Route path= "/sistema" element={<Sistema/>}/>
        <Route path= "/publications" element={<Publication/>}/>
        <Route path= "/mycomments" element={<MyComments/>}/>
      </Routes>
    </Router>
  );
}


export default App;