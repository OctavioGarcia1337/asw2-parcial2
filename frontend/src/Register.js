import React, { useState } from "react";
import "./css/Register.css";
import Cookies from "universal-cookie";

const Cookie = new Cookies();

async function register(username, password) {
  /* //NO CAMBIAR, DEBERIA ANDAR
  return await fetch('http://localhost:8090/users', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
	body: JSON.stringify({"username": username, "password":password})
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
    })*/
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

    var { uname, pass, cpass } = document.forms[0];

    // Validate user registration TODO: Check if the username is not taken
    const userData = register(uname.value, pass.value).then(data => {
      if (pass.value == cpass.value) {
        
        goto("/login")
      }
      else{
        setErrorMessages({name: "default", message: error})
      }
    })
  };


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
        <div className="button-container">
          <input type="submit"/>
        </div>
      </form>
    </div>
  );


  return (
    <div className="app">
      <div className="login-form">
        <div className="title">CREAR UN USUARIO</div>

        {isSubmitted || Cookie.get("user_id") > -1 ? Cookie.get("username") : renderForm}
      </div>
    </div>
  );
}

export default Register;