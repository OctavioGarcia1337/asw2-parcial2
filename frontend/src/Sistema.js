import React, { useState , useEffect} from "react";
import "./css/Home.css";
import logo from "./images/logo.svg"
import loadinggif from "./images/loading.gif"
import Cookies from "universal-cookie";
import {HOST, PORT} from "./config/config";
import Comments from "./Comments";


const URL = HOST + ":" + PORT
const Cookie = new Cookies();

async function getSystems(){
  return await fetch(URL + "/search=*_*", {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

async function setSystems(){
    return await fetch(URL + "/search=*_*", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      }
    }).then(response => response.json())
  }

function goto(path){
  window.location = window.location.origin + path
}

function retry() {
  goto("/")
}

function parseField(field){
  if (field !== undefined){
    return field
  }
  return "Not available"
}

function showSystem(systems){
  return systems.map((system) =>
   <div obj={system} key={system.id} className="System">
        <div>
            <a className="title">{parseField(system.titulo)}</a>
        </div>
        <div>
            <button onClick={()=>getSystems()}> GET </button>
            <button onClick={()=>setSystems()}> POST </button>
        </div>
    </div>
    ) 
}

// Probablemente hay que eliminar muchas de estas funciones
async function getSystemsBySearch(field, query){
  return fetch( URL + "/search=" + "id" + "_" + localStorage.getSystem("id"), {
    method: "GET",
    header: "Content-Type: application/json"
  }).then(response=>response.json())
}

function System() {
  const [isLogged, setIsLogged] = useState(false)
  const [user, setUser] = useState({})
  const [systems, setSystems] = useState([])
  const [needSystems, setNeedSystems] = useState(true)
  const [failedSearch, setFailedSearch] = useState(false)
  const [querying, setQuerying] = useState(false)
  const [query, setQuery] = useState("")

  if(!systems.length && needSystems){
    getSystems().then(response => setSystems(response))
    setNeedSystems(false)
  }


async function searchQuery(field, query){
    if(query == ""){
        query = localStorage.getSystem("id")
    }
    await getSystemsBySearch(field, localStorage.getSystem("id")).then(response=>{
    if(response != null){
        if(response.length > 0){
                setSystems(response)
                setFailedSearch(false)
        }else{
                setSystems([])
                setFailedSearch(true)
            }
        }
        else{
          setFailedSearch(false)
          getSystems().then(response=>setSystems(response))
        }
    })
}

  const options= (
      <div className="options-div">
        <div>
          <a onClick={()=>searchQuery("titulo", query)}>Titulo: <span>{query}</span></a>
          <a onClick={()=>searchQuery("tipo", query)}>Tipo: <span>{query}</span></a>
          <a onClick={()=>searchQuery("descripcion", query)}>Descripcion: <span>{query}</span></a>
          <a onClick={()=>searchQuery("ubicacion", query)}>Ubicacion: <span>{query}</span></a>
          <a onClick={()=>searchQuery("barrio", query)}>Barrio: <span>{query}</span></a>
          <a onClick={()=>searchQuery("vendedor", query)}>Vendedor: <span>{query}</span></a>
        </div>
      </div>
  )

  const login = (

    <span>
    <img src="./images/loading.gif" onClick={()=>goto("/user")} id="user" width="48px" height="48px"/>
    {/*<a id="logout" onClick={logout}> <span> Welcome in {user.first_name} </span> </a>*/}
    </span>
  )

  const loading = (<img id="loading" src={loadinggif}/>)

  const renderFailedSearch = (<a>No results :(</a>)

  if(query == "" && systems.length <= 0){
    searchQuery("*","*") // segundo * sacar de localstorage id
  }

  /* Funciones con cookies

  function productsByCategoryId(id, setter, categorySetter) {
    getProductsByCategoryId(id).then(response => {setter(response); 
    Cookie.set("category", id); getCategoryById(id).then(category => categorySetter(category))})
  }

  function addToCart(id, setCartSystems){
    let cookie = Cookie.get("cart");
  
    if(cookie == undefined){
      Cookie.set("cart", id + ",1;", {path: "/"});
      setCartSystems(1)
      return
    }
    let newCookie = ""
    let isNewSystem = true
    let toCompare = cookie.split(";")
    let total = 0;
    toCompare.forEach((system) => {
      if(system != ""){
        let array = system.split(",")
        let system_id = array[0]
        let system_quantity = array[1]
        if(id == system_id){
          system_quantity = Number(system_quantity) + 1
          isNewSystem = false
        }
        newCookie += system_id + "," + system_quantity + ";"
        total += Number(system_quantity);
      }
    });
    if(isNewSystem){
      newCookie += id + ",1;"
      total += 1;
    }
    cookie = newCookie
    Cookie.set("cart", cookie, {path: "/"})
    Cookie.set("cartSystems", total, {path: "/"})
    setCartSystems(total)
    return
  }*/





  return (
    <div className="home">
        <div className="topnavHOME">
            <div>
                <img src={logo} width="80px" height="80px" id="logo" onClick={()=>goto("/")} /> <p> TuCasa.com </p>
            </div>
        </div>

        <div id="mySidenav" className="sidenav" > 
          <a id="login" onClick={()=>goto("/login")}>Login</a>
          <a id="register" onClick={()=>goto("/register")}>Register</a>
          <a id="publications" onClick={()=>goto("/publications")}>Publicaciones</a>
        </div>

        <div id="main">
            {failedSearch ? renderFailedSearch : void(0)}
            {systems.length > 0 || failedSearch ? showSystem(systems) : loading}
         </div>
    </div>
    );
}

export default System;