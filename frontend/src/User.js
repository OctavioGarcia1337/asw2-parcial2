import React, { useState } from "react";
import "./css/User.css";
import logo from "./images/logo.svg"
import cart from "./images/cart.svg"
import loadinggif from "./images/loading.gif"
import usersvg from "./images/user.svg"
import Cookies from "universal-cookie";

const Cookie = new Cookies();

async function getUserById(id){
    return await fetch('http://localhost:8090/user/' + id, {
        method: 'GET',
        headers: {
        'Content-Type': 'application/json'
        }
    }).then(response => response.json())

}


function goto(path){
  window.location = window.location.origin + path
}


function logout(){
  Cookie.set("user_id", -1, {path: "/"})
  document.location.reload()
}

function User() {
  const [isLogged, setIsLogged] = useState(false)
  const [user, setUser] = useState({})


  if (Cookie.get("user_id") > -1 && !isLogged){
    getUserById(Cookie.get("user_id")).then(response => setUser(response))
    setIsLogged(true)
  }

  if (!(Cookie.get("user_id") > -1) && isLogged){
    setIsLogged(false)
  }

    const login = (

    <span>
    <img src={usersvg} onClick={()=>goto("/user")} id="user" width="48px" height="48px"/>
    <img src={cart} onClick={()=>goto("/cart ")} id="cart" width="48px" height="48px"/>
    <a id="logout" onClick={logout}> <span> Welcome in {user.first_name} </span> </a>
    </span>
  )

  const showUserInfo = (
    <div className="userInfo">
      <img src={usersvg} width="128px" height="128px"/>
      <div> {user.first_name} {user.last_name} </div>
      <div> Username: {user.username} </div>
      <div> {user.first_name}, {user.last_name} </div>
      <div> Email: {user.email} </div>
    </div>
  )

  const pleaseLogin = (
    <div> Nothing to show. Please login and maybe we'll get some info for ya </div>
  )


  const loading = (
    <img id="loading" src={loadinggif}/>
  )

  return (
    <div className="home">
      <div className="topnav">
            <div>
                <img src={logo} width="80px" height="80px" id="logo" onClick={()=>goto("/")} /> <p> TuCasa.com </p>
            </div>
        {isLogged ? login : <a id="login" onClick={()=>goto("/login")}>Login</a>}
      </div>

      <div id="main">
        {isLogged ? showUserInfo : pleaseLogin}
      </div>
    </div>
  );
}

export default User;