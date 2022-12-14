import React, { useState } from "react";
import "./css/Orders.css";
import logo from "./images/logo.svg"
import cart from "./images/cart.svg"
import usersvg from "./images/user.svg"
import Cookies from "universal-cookie";

const Cookie = new Cookies();

async function getUserById(id){
  return fetch("http://localhost:8090/user/" + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

async function getitemById(id){
  return fetch("http://localhost:8090/item/" + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

async function getOrderById(id) {
  return fetch("http://localhost:8090/publications/" + id, {
    method:"GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

async function deletePublication(id) {
    return fetch("http://localhost:8090/publications/" + id, {
      method:"DELETE",
      headers: {
        "Content-Type": "application/json"
      }
    }).then(response => response.json())
  }

function parseField(field){
    if (field !== undefined){
      return field
    }
    return "Not available"
  }

function goto(path){
  window.location = window.location.origin + path
}

function logout(){
  Cookie.set("user_id", -1, {path:"/"})
  window.location.reload();
}

function showItems(items){
  return items.map((item) =>
  <div obj={item} key={item.id} className="item" onClick={()=>test(item.id)}>
  <div>
    <button onClick={()=>deletePublication(item.id)}> Eliminar </button>
    <img width="128px" height="128px" src={parseField(item.url_img)}  onError={(e) => (e.target.onerror = null, e.target.src = "./images/default.jpg")}/>
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
  <a className="comentarios"> Ver Comentarios</a>
  </div>
 </div>
 )

}

function showAddress(address){
  return (
    <div className="orderAddress">
      ADDRESS N°{address.address_id}
      <div><span className="orderAddressInfo"> Street: </span> <a className="orderAddressInfoLoad">{address.street1}</a> </div>
      <div><span className="orderAddressInfo"> Street2: </span> <a className="orderAddressInfoLoad">{address.street2} </a> </div>
      <div><span className="orderAddressInfo"> Number: </span> <a className="orderAddressInfoLoad">{address.number} </a> </div>
      <div><span className="orderAddressInfo"> District: </span> <a className="orderAddressInfoLoad">{address.district} </a> </div>
      <div><span className="orderAddressInfo"> City: </span> <a className="orderAddressInfoLoad">{address.city} </a> </div>
      <div><span className="orderAddressInfo"> Country: </span> <a className="orderAddressInfoLoad">{address.country} </a> </div>
    </div>
  )

}


async function getOrderitems(){
  let items = []
  let a = Cookie.get("order").split(";")

  for (let i = 0; i < a.length; i++){
    let item = a[i];
    if(item != ""){
      let array = item.split(",")
      let id = array[0]
      let quantity = array[1]
      let item = await getitemById(id)
      item.quantity = quantity;
      items.push(item)
    }
  }
  return items
}

async function getAddressById(id){
  return fetch("http://localhost:8090/address/" + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}


async function setOrder(setOrder, setTotal){
  let total = 0;
  await getOrderitems().then(response => {
    setOrder(response)
    response.forEach((item) => {
      total += item.base_price * item.quantity;
    });
    setTotal(total)
  })
}



function Publication(){
  const [user, setUser] = useState({});
  const [isLogged, setIsLogged] = useState(false);
  const [orderitems, setOrderitems] = useState([])
  const [total, setTotal] = useState(0)
  const [address, setAddress] = useState({})

  const login = (

    <span>
    <img src={usersvg} onClick={()=>goto("/user")} id="user" width="48px" height="48px"/>
    <img src={cart} onClick={()=>goto("/cart")} id="cart" width="48px" height="48px"/>
    <a id="logout" onClick={logout}> <span> Welcome in {user.first_name} </span> </a>
    </span>
  )

  if (Cookie.get("user_id") > -1 && !isLogged) {
    getUserById(Cookie.get("user_id")).then(response => setUser(response))
    setIsLogged(true)

    if (Cookie.get("address") > -1){
      getAddressById(Cookie.get("address")).then(response => setAddress(response))
    }
  }


  if (orderitems.length <= 0 && Cookie.get("user_id") > -1){
    setOrder(setOrderitems, setTotal)
  }

  const complete = (
    <div>
    <div> Woohoo you placed a freaking order. I'm so proud of you</div>
    {address.address_id != undefined ? showAddress(address) : <span> NOOO </span>}
    {showItems(orderitems)}


    <div> Total: ${total} </div>
    </div>
  )


  const error = (
    <div>
    <div> BOO ERROR :(((( </div>
    <div> Let's think. This probably happened because of some stock mistake </div>
    <div> Error {Cookie.get("orderError")} </div>
    </div>
  )

  return (
    <div className="orderstatus">
      <div className="topnav">
        <img src={logo} width="80px" height="80px" id="logo" onClick={()=>goto("/")} />
      </div>

      <div id="mySidenav" className="sidenav" > 
        <a id="login" onClick={()=>goto("/login")}>Login</a>
        <a id="register" onClick={()=>goto("/register")}>Register</a>
        <a id="sistema" onClick={()=>goto("/sistema")}>Sistema</a>
      </div>

      <div id="main">
        {window.location.pathname.split("/")[2] == "complete" ? complete : error}
      </div>
    </div>
  );
}

export default Publication;