import React, { useState , useEffect} from "react";
import "./css/Item.css";
import logo from "./images/logo.svg"
import loadinggif from "./images/loading.gif"
import Cookies from "universal-cookie";
import {HOST, PORT, ITEMSPORT} from "./config/config";
import Comments from "./Comments";


const URL = HOST + ":" + PORT
const ITEMSURL = HOST + ":" + ITEMSPORT
const Cookie = new Cookies();

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

function showItem(item){
  return (
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

async function getItemById(id){
    return fetch(ITEMSURL + "/items/" + id, {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        }
    }).then(response => response.json())
}

function Item() {
  const [isLogged, setIsLogged] = useState(false)
  const [user, setUser] = useState({id: null})
  const [needItem, setNeedItem] = useState(true)
  const [item, setItem] = useState({})
  const [failedSearch, setFailedSearch] = useState(false)

    if (needItem){
        Cookie.set("need_item", "true");
        setNeedItem(false);
    }


  const login = (

    <span>
    <img src="./images/loading.gif" onClick={()=>goto("/user")} id="user" width="48px" height="48px"/>
    {/*<a id="logout" onClick={logout}> <span> Welcome in {user.first_name} </span> </a>*/}
    </span>
  )

  const loading = (<img id="loading" src={loadinggif}/>)

  const renderFailedSearch = (<a>No results :(</a>)

    if (Cookie.get("need_item") === "true") {
        getItemById(localStorage.getItem("id")).then(response => setItem(response));
        Cookie.set("need_item", "false");
    }

    /* Funciones con cookies
    function productsByCategoryId(id, setter, categorySetter) {
      getProductsByCategoryId(id).then(response => {setter(response);
      Cookie.set("category", id); getCategoryById(id).then(category => categorySetter(category))})
    }
    function addToCart(id, setCartItems){
      let cookie = Cookie.get("cart");
      if(cookie == undefined){
        Cookie.set("cart", id + ",1;", {path: "/"});
        setCartItems(1)
        return
      }
      let newCookie = ""
      let isNewItem = true
      let toCompare = cookie.split(";")
      let total = 0;
      toCompare.forEach((item) => {
        if(item != ""){
          let array = item.split(",")
          let item_id = array[0]
          let item_quantity = array[1]
          if(id == item_id){
            item_quantity = Number(item_quantity) + 1
            isNewItem = false
          }
          newCookie += item_id + "," + item_quantity + ";"
          total += Number(item_quantity);
        }
      });
      if(isNewItem){
        newCookie += id + ",1;"
        total += 1;
      }
      cookie = newCookie
      Cookie.set("cart", cookie, {path: "/"})
      Cookie.set("cartItems", total, {path: "/"})
      setCartItems(total)
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
          <a id="sistema" onClick={()=>goto("/sistema")}>Sistema</a>
          <a id="publications" onClick={()=>goto("/publications")}>Publicaciones</a>
        </div>

        <div id="main">
            {failedSearch ? renderFailedSearch : void(0)}
            {item.id != null ? showItem(item) : loading}
         </div>
    </div>
    );
}

export default Item;