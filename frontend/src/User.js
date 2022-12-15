import React, { useState } from "react";
import "./css/User.css";
import logo from "./images/logo.svg"
import cart from "./images/cart.svg"
import loadinggif from "./images/loading.gif"
import usersvg from "./images/user.svg"
import Cookies from "universal-cookie";
import { HOST, PORT, USERSPORT } from "./config/config";

const Cookie = new Cookies();
const URL = `${HOST}:${USERSPORT}`

async function getUserById(id) {
  return await fetch(`${URL}/users/${id}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'user/json'
    }
  }).then(response => response.json())

}


function goto(path) {
  window.location = window.location.origin + path
}

async function deleteUser(id) {
  return await fetch(`${URL}/user/${id}`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => {
    response.status === 200 ? goto("/") : alert("error deleting account");
  })
}


function logout() {
  Cookie.set("user_id", -1, { path: "/" })
  document.location.reload()
}

function User() {
  const [isLogged, setIsLogged] = useState(false)
  const [user, setUser] = useState({})


  if (Cookie.get("user_id") > -1 && !isLogged) {
    getUserById(Cookie.get("user_id")).then(response => setUser(response))
    setIsLogged(true)
  }

  if (!(Cookie.get("user_id") > -1) && isLogged) {
    setIsLogged(false)
  }

  const login = (
    <span>
      <img src={usersvg} onClick={() => goto("/user")} id="user" width="48px" height="48px" />
    </span>
  )

  const showUserInfo = (
    <div>
      <div className="userInfo">
        <img src={usersvg} width="128px" height="128px" />
        <div> {user.first_name} {user.last_name} </div>
        <div> Username: {user.username} </div>
        <div> {user.first_name}, {user.last_name} </div>
        <div> Email: {user.email} </div>
      </div>
      <div id="eliminar">
        <button id="eliminar-cuenta" onClick={() => deleteUser(user.user_id)}> Eliminar Cuenta </button>
      </div>
    </div>
  )

  const pleaseLogin = (
    <div> Nothing to show. Please login and maybe we'll get some info for ya </div>
  )


  const loading = (
    <img id="loading" src={loadinggif} />
  )

  const logreg = (
    <div>
      <a id="login" onClick={() => goto("/login")}>Login</a>
      <a id="register" onClick={() => goto("/register")}>Register</a>
    </div>
  )

  const loggedout = (
    <div>
      <a id="logout" onClick={logout}> <span> Welcome in {user.first_name} </span> </a>
    </div>
  )

  return (
    <div className="home">
      <div className="topnavHOME">
        <div>
          <img src={logo} width="80px" height="80px" id="logo" onClick={() => goto("/")} /> <p> TuCasa.com </p>
          {isLogged ? login : <a id="login" onClick={() => goto("/login")}>Login</a>}
        </div>
      </div>  
      <div>
        <div id="mySidenav" className="sidenav" >
          {isLogged ? loggedout : logreg}
          <a id="sistema" className="clicked" onClick={() => goto("/sistema")}>Sistema</a>
          <a id="publications" onClick={() => goto("/publications")}>Publicaciones</a>
          <a id="mycomments" onClick={() => goto("/mycomments")}>Mis Comentarios</a>
        </div>
      </div>

      <div id="main">
        {isLogged ? showUserInfo : pleaseLogin}
      </div>
  </div>
  );
}

export default User;