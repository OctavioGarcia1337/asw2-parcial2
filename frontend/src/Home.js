import React, { useState , useEffect} from "react";
import "./css/Home.css";
import logo from "./images/logo.svg"
import loadinggif from "./images/loading.gif"
import usersvg from "./images/user.svg"
import Cookies from "universal-cookie";
import {HOST, PORT} from "./config/config";


const URL = HOST + ":" + PORT
const Cookie = new Cookies();

async function getItems(){
  return await fetch(URL + "/search=*_*", {
    method: "GET",
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


function showItems(items){
  return items.map((item) =>

   <div obj={item} key={item.id} className="item">
    <div onClick={()=>goto("/item?id="+item.id)}>
      <img width="128px" height="128px" src={item.Url_Img}  onError={(e) => (e.target.onerror = null, e.target.src = "./images/default.jpg")}/>
    </div>
    <a className="name">{item.Titulo}</a>
    <a className="price"> - {"$" + item.Precio_base}</a>
    <a className="price"> -  Expensas: {"$" + item.Expensas}</a>
    <div>
      <a className="description">{item.Tipo}</a>
    </div>
    <div>
      <a className="description">{item.Ubicacion}</a>
      <a className="description">, {item.Barrio}</a>
    </div>
    <div>
      <a className="description">{item.Descripcion}</a>
    </div>
    <div>
      <a className="description">{item.Vendedor}</a>
    </div>
    <div className="right">
      <a className="stock">Mts2: {item.Mts2}</a>
      <a className="stock"> - Ambientes: {item.Ambientes}</a>
      <a className="stock"> - Dormitorios: {item.Dormitorios}</a>
      <a className="stock"> - Baños: {item.Banos}</a>
    </div>
   </div>
 )//agregar los campos faltantes
}


async function getItemsBySearch(field, query){
  return fetch( URL + "/search=" + field + "_" + query, {
    method: "GET",
    header: "Content-Type: application/json"
  }).then(response=>response.json())
}

async function getItemsBySearchAll(query){
  return fetch( URL + "/searchAll=" + query, {
    method: "GET",
    header: "Content-Type: application/json"
  }).then(response=>response.json())
}

function Home() {
  const [isLogged, setIsLogged] = useState(false)
  const [user, setUser] = useState({})
  const [items, setItems] = useState([])
  const [needItems, setNeedItems] = useState(true)
  const [failedSearch, setFailedSearch] = useState(false)

  if(!items.length && needItems){
    getItems().then(response => setItems(response))
    setNeedItems(false)
  }

  async function searchQueryAll(query){

    await getItemsBySearchAll(query).then(response=>{
      if(query == ""){
        query = "*"
      }
      if(response != null){
        if(response.length > 0){
          setItems(response)
          setFailedSearch(false)
        }else{
          setItems([])
          setFailedSearch(true)
        }
      }
      else{
        setFailedSearch(false)
        getItems().then(response=>setItems(response))
      }
    })


  }

  const login = (

    <span>
    <img src={usersvg} onClick={()=>goto("/user")} id="user" width="48px" height="48px"/>
    {/*<a id="logout" onClick={logout}> <span> Welcome in {user.first_name} </span> </a>*/}
    </span>
  )

  const loading = (
    <img id="loading" src={loadinggif}/>
  )

  const renderFailedSearch = (
    <a>No results :(</a>
  )

  return (
    <div className="home">
      <div className="topnavHOME">
        <div>
          <img src={logo} width="80px" height="80px" id="logo" onClick={()=>goto("/")} /> <p>3 Random Words Shop</p>
        </div>
        <input type="text" id="search" placeholder="Search..." onKeyDown={(e) => e.keyCode === 13 ? searchQueryAll(e.target.value) : void(0)}/>
        {isLogged ? login : <a id="login" onClick={()=>goto("/login")}>Login</a>}
      </div>


      <div id="mySidenav" className="sidenav">

      </div>

      <div id="main">
        {failedSearch ? renderFailedSearch : void(0)}
        {items.length > 0 || failedSearch ? showItems(items) : loading}

      </div>
    </div>
  );
}

export default Home;
