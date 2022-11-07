import React, { useState , useEffect} from "react";
import "./css/Home.css";
import logo from "./images/logo.svg"
import loadinggif from "./images/loading.gif"
import usersvg from "./images/user.svg"
import Cookies from "universal-cookie";

const Cookie = new Cookies();

/* 
async function getUserById(id){
    return await fetch('http://localhost:8090/user/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
}).then(response => response.json())

}*/

async function getProducts(){
  return await fetch('http://localhost:8090/products', {
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


function showProducts(products, setCartItems){
  return products.map((product) =>

   <div obj={product} key={product.product_id} className="product">
    <div onClick={()=>goto("/product?id="+product.product_id)}>
      <img width="128px" height="128px" src={"./images/" + product.picture_url}  onError={(e) => (e.target.onerror = null, e.target.src = "./images/default.jpg")}/>
    </div>
    <a className="name">{product.name}</a>
    {/* <a className="addcart" onClick={() => addToCart(product.product_id, setCartItems)}>Add to Cart</a> */}
    <a className="price">{product.currency_id + "$" + product.base_price}</a>
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


function search(){
  let input, filter, a, i;
  input = document.getElementById("search");
  filter = input.value.toUpperCase();
  a = document.getElementsByClassName("product");
  for (i = 0; i < a.length; i++) {
    let txtValue = a[i].children[1].textContent || a[i].children[1].innerText;
    if (txtValue.toUpperCase().indexOf(filter) > -1) {
      a[i].style.display = "inherit";
    } else {
      a[i].style.display = "none";
    }
  }
  if(input.value.toUpperCase().length <= 0){
    for(i = 0; i < a.length; i++){
      a[i].style.display = "inherit";
    }
  }

}

function deleteCategory(){
  Cookie.set("category", 0, {path: "/"})
  goto("/")
}

async function getProductBySearch(query){
  return fetch("http://localhost:8090/products/search="+query, {
    method: "GET",
    header: "Content-Type: application/json"
  }).then(response=>response.json())
}


function Home() {
  const [isLogged, setIsLogged] = useState(false)
  const [user, setUser] = useState({})
  const [categories, setCategories] = useState([])
  const [products, setProducts] = useState([])
  const [needProducts, setNeedProducts] = useState(true)
  const [category, setCategory] = useState("")
  const [needCategories, setNeedCategories] = useState(true)
  const [cartItems, setCartItems] = useState("")
  const [failedSearch, setFailedSearch] = useState(false)

/*
  if (Cookie.get("user_id") > -1 && !isLogged){
    getUserById(Cookie.get("user_id")).then(response => setUser(response))
    setIsLogged(true)
  }

  if (!(Cookie.get("user_id") > -1) && isLogged){
    setIsLogged(false)
  }

  if(!categories.length && needCategories){
    getCategories().then(response => setCategories(response))
    setNeedCategories(false)
  }*/

  if(!products.length && needProducts){
    getProducts().then(response => setProducts(response))
    setNeedProducts(false)
  }

/*
  if (!cartItems && Cookie.get("cartItems")){
    setCartItems(Cookie.get("cartItems"))
  }*/


  async function searchQuery(query){

    await getProductBySearch(query).then(response=>{
      console.log(query)
      if(response != null){
        if(response.length > 0){
          setProducts(response)
          setFailedSearch(false)
        }else{
          setProducts([])
          setFailedSearch(true)
        }
      }
      else{
        setFailedSearch(false)
        getProducts().then(response=>setProducts(response))
      }
    })


  }

  const login = (

    <span>
    <img src={usersvg} onClick={()=>goto("/user")} id="user" width="48px" height="48px"/>
    {/* <img src={cart} onClick={()=>goto("/cart ")} id="cart" width="48px" height="48px"/> */}
    <span className="cartNumber">{cartItems > 0 ? cartItems : 0}</span>
    {/* <a id="logout" onClick={logout}> <span> Welcome in {user.first_name} </span> </a> */}
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
        <input type="text" id="search" placeholder="Search..." onKeyDown={(e) => e.keyCode === 13 ? searchQuery(e.target.value) : void(0)}/>
        {isLogged ? login : <a id="login" onClick={()=>goto("/login")}>Login</a>}
      </div>


      <div id="mySidenav" className="sidenav">

         {/* {categories.length > 0 ? showCategories(categories, setProducts, setCategory) : <a onClick={retry}> Loading Failed. Click to retry </a>} */}
      </div>

      <div id="main">
        {failedSearch ? renderFailedSearch : void(0)}
        {Cookie.get("category") > 0 ? <a className="categoryFilter"> {category.name} <button className="delete" onClick={deleteCategory}>X</button> </a> : <a/>}
        {products.length > 0 || failedSearch ? showProducts(products, setCartItems) : loading}


      </div>
    </div>
  );
}

export default Home;
