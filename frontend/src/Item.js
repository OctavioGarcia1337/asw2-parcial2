import React, { useState , useEffect} from "react";
import "./css/Item.css";
import logo from "./images/logo.svg"
import loadinggif from "./images/loading.gif"
import Cookies from "universal-cookie";
import {HOST, PORT} from "./config/config";
import Comments from "./Comments";


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

function parseField(field){
  if (field !== undefined){
    return field
  }
  return "Not available"
}



function showItem(items){
  return items.map((item) =>
   <div obj={item} key={item.id} className="item">
        <div>
            <img width="240px" height="240px" src={parseField(item.url_img)}  onError={(e) => (e.target.onerror = null, e.target.src = "./images/default.jpg")}/>
        </div>
            <a className="title">{parseField(item.titulo)}</a>
            <a className="price"> {"$" + parseField(item.precio_base)}</a>
        <div>
            <a className="expenses"> -  Expensas: {"$" + parseField(item.expensas)}</a>
        </div>
        <div>
            <a className="type">{parseField(item.tipo)}</a>
        </div>
        <div>
            <a className="location">{parseField(item.ubicacion)},</a>
            <a className="neighbourhood">{parseField(item.barrio)}</a>
        </div>
        <div>
            <a className="description">{parseField(item.descripcion)}</a>
        </div>
        <div className="sellerBlock">
            <a className="seller">{parseField(item.vendedor)}</a>
        </div>
        <div className="right">
            <a className="sqmts">Mts2: {parseField(item.mts2)}</a>
            <a className="rooms"> - Ambientes: {parseField(item.ambientes)}</a>
            <a className="bedrooms"> - Dormitorios: {parseField(item.dormitorios)}</a>
            <a className="bathrooms"> - Baños: {parseField(item.banos)}</a>
        </div>
        <div>
          <Comments CurrentUserId="1" />
        </div>
    </div>
    ) 
}


async function getItemsBySearch(field, query){
  return fetch( URL + "/search=" + "id" + "_" + localStorage.getItem("id"), {
    method: "GET",
    header: "Content-Type: application/json"
  }).then(response=>response.json())
}

function Item() {
  const [items, setItems] = useState([])
  const [needItems, setNeedItems] = useState(true)
  const [failedSearch, setFailedSearch] = useState(false)
  const [querying, setQuerying] = useState(false)
  const [query, setQuery] = useState("")

  if(!items.length && needItems){
    getItems().then(response => setItems(response))
    setNeedItems(false)
  }


async function searchQuery(field, query){
    if(query == ""){
        query = localStorage.getItem("id")
    }
    await getItemsBySearch(field, localStorage.getItem("id")).then(response=>{
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

  const options= (
      <div className="options-div">
        <div>
          <a onClick={()=>searchQuery("titulo", query)}>Titulo: <span>{query}</span></a>
          <a onClick={()=>searchQuery("titulo", query)}>Tipo: <span>{query}</span></a>
          <a onClick={()=>searchQuery("titulo", query)}>Descripcion: <span>{query}</span></a>
          <a onClick={()=>searchQuery("titulo", query)}>Ubicacion: <span>{query}</span></a>
          <a onClick={()=>searchQuery("titulo", query)}>Barrio: <span>{query}</span></a>
          <a onClick={()=>searchQuery("titulo", query)}>Vendedor: <span>{query}</span></a>
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

  if(query == "" && items.length <= 0){
    searchQuery("*","*") // segundo * sacar de localstorage id
  }

  return (
    <div className="home">
        <div className="topnavHOME">
            <div>
                <img src={logo} width="80px" height="80px" id="logo" onClick={()=>goto("/")} /> <p> TuCasa.com </p>
            </div>
        </div>

        <div id="mySidenav" className="sidenav"></div>

        <div id="main">
            {failedSearch ? renderFailedSearch : void(0)}
            {items.length > 0 || failedSearch ? showItem(items) : loading}
         </div>
    </div>
    );
}

export default Item;