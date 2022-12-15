import React, { useState } from "react";
import "./css/Register.css";
import Cookies from "universal-cookie";
import logo from "./images/logo.svg";
import usersvg from "./images/user.svg";
import {HOST, PORT, USERSHOST,USERSPORT} from "./config/config";

const Cookie = new Cookies();
const URL = `${USERSHOST}:${USERSPORT}`

async function register(username, password, first_name, last_name, email) {
  return await fetch(`${URL}/user`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
	body: JSON.stringify({
      "username": username,
      "password":password,
      "first_name":first_name,
      "last_name":last_name,
      "email":email
    })
  })
    .then(response => {
      if (response.status == 400 || response.status == 401)
      {
        return {"user_id": -1}
      }
      return response.json()
    })
    .then(response => {
      Cookie.set("user_id", response.user_id, {path: '/'})
      Cookie.set("username", username, {path: '/login'})
    })
}

function goto(path){
  window.location = window.location.origin + path
}

function Register() {

  // React States
  const [errorMessages, setErrorMessages] = useState({});
  const [isSubmitted, setIsSubmitted] = useState(false);

  const error = "Las contraseñas deben coincidir";

  const handleSubmit = (event) => {
    //Prevent page reload
    event.preventDefault();

    var { uname, pass, cpass, first_name, last_name, email } = document.forms[0];

    // Validate user registration TODO: Check if the username is not taken
    register(uname.value, pass.value, first_name.value, last_name.value, email.value).then(data => {
      if (pass.value === cpass.value) {
        goto("/login")
      }
      else{
        setErrorMessages({name: "default", message: error})
      }
    })
  };

  const login = (
    <span>
        <img src={usersvg} onClick={()=>goto("/user")} id="user" width="48px" height="48px"/>
    </span>
  )

  // Generate JSX code for error message
  const renderErrorMessage = (name) =>
    name === errorMessages.name && (
      <div className="error">{errorMessages.message}</div>
    );

  // JSX code for login form
  const renderForm = (
    <div className="form">
      <form onSubmit={handleSubmit}>
        <div className="input-container">
          <label>Usuario </label>
          <input type="text" name="uname" placeholder="Usuario" required />
        </div>
        <div className="input-container">
          <label>Contraseña</label>
          <input type="password" name="pass" placeholder="Contraseña" required />
        </div>
        <div className="input-container">
          <label>Confirmar Contraseña</label>
          <input type="password" name="cpass" placeholder="Repetir Contraseña" required />
        </div>

          {renderErrorMessage("default")}

        <div className="input-container">
          <label>Nombre</label>
          <input type="text" name="first_name" placeholder="Nombre" required />
        </div>
        <div className="input-container">
          <label>Apellido</label>
          <input type="text" name="last_name" placeholder="Apellido" required />
        </div>
        <div className="input-container">
          <label>Email</label>
          <input type="email" name="email" placeholder="EMAIL@HOST.COM" required />
        </div>
        <div className="button-container">
          <input type="submit"/>
        </div>
      </form>
    </div>
  );

  const logreg = (
      <div>
        <a id="login" onClick={()=>goto("/login")}>Login</a>
        <a id="register" className="clicked" onClick={()=>goto("/register")}>Register</a>
      </div>
  )

  return (
    <div>
    <div className="home">
      <div className="topnavHOME">
        <div>
          <img src={logo} width="80px" height="80px" id="logo" onClick={()=>goto("/")} /> <p> TuCasa.com </p>
        </div>
      </div>
    </div>

    <div id="mySidenav" className="sidenav" >
      {logreg}
      <a id="sistema" onClick={()=>goto("/sistema")}>Sistema</a>
      <a id="publications" onClick={()=>goto("/publications")}>Publicaciones</a>
      <a id="mycomments" onClick={()=>goto("/mycomments")}>Mis Comentarios</a>
    </div>

    <div className="app">
      <div className="login-form">
        <div className="title">CREAR UN USUARIO</div>

        {isSubmitted || Cookie.get("user_id") > -1 ? Cookie.get("username") : renderForm}
      </div>
    </div>
    </div>
  );
}

export default Register;