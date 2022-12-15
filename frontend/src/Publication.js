import React, {useEffect, useState} from "react";
import "./css/Orders.css";
import logo from "./images/logo.svg"
import cart from "./images/cart.svg"
import usersvg from "./images/user.svg"
import Cookies from "universal-cookie";
import "./css/Item.css";
import PublicationForm from "./PublicationForm";
import {HOST, ITEMSPORT, USERSPORT} from "./config/config";

const Cookie = new Cookies();
const URLITEMS = `${HOST}:${ITEMSPORT}`
const URLUSERS = `${HOST}:${USERSPORT}`

function logout(){
    Cookie.set("user_id", -1, {path: "/"})
    document.location.reload()
}

async function getUserById(id) {
    return fetch(`${URLUSERS}/users/` + id, {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        }
    }).then(response => response.json())
}

function goToItem(id){
    window.localStorage.setItem("id",id)
    goto("/item")
}


async function getItemById(id) {
    return fetch(`${URLITEMS}/item/` + id, {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        }
    }).then(response => response.json())
}

async function makeItem(newItem) {
    let userid = Number(Cookie.get("user_id"))
    let parsedItem = JSON.parse(newItem)
    if (parsedItem.length) {
        for (let i = 0; i < parsedItem.length; i++){
            parsedItem[i].usuario_id = userid;
        }
    }
    else{
        parsedItem.usuario_id = userid;
        parsedItem = [parsedItem]
    }

    return await fetch(`${URLITEMS}/items`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(parsedItem)
    })
      .then(response => {
        if (response.status == 400 || response.status == 401)
        {
          return {"json_err": -1}
        }
        return response.json()
      })
  }

async function deleteItem(id) {
    return await fetch(`${URLITEMS}/item/` + id, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json"
        }
    }).then(response => {
        response.status === 200 ? goto("/publications") : alert("error deleting item");
    })
}

function parseField(field) {
    if (field !== undefined) {
        return field
    }
    return "Not available"
}

function goto(path) {
    window.location = window.location.origin + path
}


function showItems(items) {


    return items.map((item) =>
        <div>
        <div obj={item} key={item.id} className="item" onClick={() => goToItem(item.id)}>
            <div>
                <img width="128px" height="128px" src={parseField(item.url_img)} onError={(e) => (e.target.onerror = null, e.target.src = "./images/default.jpg")} />
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


        </div>
            <div id="eliminar">
                <button id="eliminar-boton" onClick={() => deleteItem(item.id)}> X </button>
            </div>
        </div>
    )

}


async function getItemsByUserId(id) {
    return fetch(`${URLITEMS}/users/${id}/items`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        }
    }).then(response => response.json())
}



async function setItems(setUserItems, userId) {
    await getItemsByUserId(userId).then(response => {
        response != null ? setUserItems(response) : setUserItems([]);
    })
}



function Item() {
    const [user, setUser] = useState({});
    const [isLogged, setIsLogged] = useState(false);
    const [userItems, setUserItems] = useState([])
    const [needItems, setNeedItems] = useState(true)
    if (Cookie.get("user_id") > -1 && !isLogged) {
        getUserById(Cookie.get("user_id")).then(response => setUser(response))
        setIsLogged(true)

    }

    useEffect(() => {
        if (userItems.length <= 0 && Cookie.get("user_id") > -1 && needItems) {
            setItems(setUserItems, Cookie.get("user_id"))
            setNeedItems(false)
        }
    }, [userItems.length])

    
    const login = (
        <span>
            <img src={usersvg} onClick={()=>goto("/user")} id="user" width="48px" height="48px"/>
        </span>
    )



    const addItem = (text) => {makeItem(text);};

    const error = (
        <div>
            <div> BOO ERROR :(((( </div>
            <div> There's no items yet :D </div>
        </div>
    )

    const logreg = (
        <div>
            <a id="login" onClick={()=>goto("/login")}>Login</a>
            <a id="register" onClick={()=>goto("/register")}>Register</a>
        </div>
    )

    const loggedout = (
        <div>
            <a id="logout" onClick={logout}> <span> Welcome in {user.first_name} </span> </a>
        </div>
    )

    return (
        <div className="items">
            <div className="topnavHOME">
                <img src={logo} width="80px" height="80px" id="logo" onClick={()=>goto("/")} /> <p> TuCasa.com </p>
                {isLogged ? login : <a id="login" onClick={()=>goto("/login")}>Login</a>}
            </div>

            <div id="mySidenav" className="sidenav" >
                {isLogged ? loggedout : logreg}
                <a id="sistema" onClick={()=>goto("/sistema")}>Sistema</a>
                <a id="publications" className="clicked" onClick={()=>goto("/publications")}>Publicaciones</a>
                <a id="mycomments" onClick={()=>goto("/mycomments")}>Mis Comentarios</a>
            </div>

            <div id="main">



                <div className="comments">
                    <h3 className="comments-title">Nueva Publicación</h3>
                    <div className="comment-form-title">JSON Here</div>
                    <PublicationForm submitLabel="Write" handleSubmit={addItem} />
                </div>

                {userItems.length > 0 ? showItems(userItems) : error}
            </div>
        </div>
    );
}

export default Item;