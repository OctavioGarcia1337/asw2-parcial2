import React, { useState } from "react";
import "./css/Product.css";
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

async function getProductById(id){
  return fetch("http://localhost:8090/product/" + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

async function getCategories(){
  return await fetch('http://localhost:8090/categories', {
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

function showCategories(categories) {
  return categories.map((category, i) => <a onClick={() => goto("/")} obj={category} key={category.category_id}>{category.name}</a>)
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
}


function logout(){
  Cookie.set("user_id", -1, {path: "/"})
  document.location.reload()
}
function showProduct(product, setCartItems){
    return(
    <div obj={product} key={product.product_id} className="productJS">
       <div>
         <img width="400px" height="400px" src={"./images/" + product.picture_url}  onError={(e) => (e.target.onerror = null, e.target.src = "./images/default.jpg")}/>
       </div>
       <div className="top">
       <a className="name">{product.name}</a>
       <a className="addcart" onClick={() => addToCart(product.product_id, setCartItems)}>Add to Cart</a>
       <a className="price">{product.currency_id + "$" + product.base_price}</a>
       </div>
       <div>
         <a className="description">{product.description}</a>
       </div>
       <div className="right">
         <a className="category">{product.category.name}</a>
         <a className="stock">Stock: {product.stock}</a>
       </div>
    </div>
  )
}

function Product() {
  const [isLogged, setIsLogged] = useState(false)
  const [user, setUser] = useState({})
  const [categories, setCategories] = useState([])
  const [product, setProduct] = useState({})
  const [category, setCategory] = useState("")
  const [cartItems, setCartItems] = useState("")


  if (Cookie.get("user_id") > -1 && !isLogged){
    getUserById(Cookie.get("user_id")).then(response => setUser(response))
    setIsLogged(true)
  }

  if (!(Cookie.get("user_id") > -1) && isLogged){
    setIsLogged(false)
  }

  if(categories.length <= 0){
    getCategories().then(response => setCategories(response))
  }

  if (!cartItems && Cookie.get("cartItems")){
    setCartItems(Cookie.get("cartItems"))
  }

  if (!(product.product_id > -1)) {
    let id = window.location.search.split("=")[1]
    getProductById(Number(id)).then(response => {setProduct(response);})
  }

  const login = (
    <span>
    <img src={usersvg} onClick={()=>goto("/user")} id="user" width="48px" height="48px"/>
    <img src={cart} onClick={()=>goto("/cart ")} id="cart" width="48px" height="48px"/>
    <span className="cartNumber">{cartItems > 0 ? cartItems : 0}</span>
    <a id="logout" onClick={logout}> <span> Welcome in {user.first_name} </span> </a>
    </span>
  )
  const loading = (
    <img id="loading" src={loadinggif}/>
  )



  return (
    <div className="home">
      <div className="topnav">
        <div>
          <img src={logo} width="80px" height="80px" id="logo" onClick={()=>goto("/")} /> <p>3 Random Words Shop</p>
        </div>
        {isLogged ? login : <a id="login" onClick={()=>goto("/login")}>Login</a>}
      </div>


      <div id="mySidenav" className="sidenav">

        {categories.length > 0 ? showCategories(categories) : <a onClick={retry}> Loading Failed. Click to retry </a>}
      </div>

      <div id="main">
        {product.product_id > -1 ? showProduct(product, setCartItems) : <a> Bad Request </a>}

      </div>
    </div>
  );
}

export default Product;
