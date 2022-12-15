import React, { useState } from "react";
import "./css/Orders.css";
import logo from "./images/logo.svg"
import Cookies from "universal-cookie";
import "./css/Home.css";
import { HOST, ITEMSPORT, USERSPORT } from "./config/config";
import Comment from "./Comment"

const Cookie = new Cookies();
const URLITEMS = `${HOST}:${ITEMSPORT}`
const URLUSERS = `${HOST}:${USERSPORT}`

function logout() {
    Cookie.set("user_id", -1, { path: "/" })
    document.location.reload()
}

async function getUserById(id) {
    return fetch(`${URLUSERS}/users/` + id, {
        method: "GET",
        headers: {
            "Content-Type": "comment/json"
        }
    }).then(response => response.json())
}

function goToComment(id) {
    window.localStorage.setComment("id", id)
    goto("/comment")
}


async function getCommentById(id) {
    return fetch(`${URLITEMS}/comment/` + id, {
        method: "GET",
        headers: {
            "Content-Type": "comment/json"
        }
    }).then(response => response.json())
}

async function deleteComment(id) {
    return await fetch(`${URLITEMS}/comment/` + id, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json"
        }
    }).then(response => {
        response.status === 200 ? goto("/comments") : alert("error deleting comment");
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


function showComments(comments) {

    return comments.map((comment) =>
        <div>
            <div>
                <Comment/>
            </div> 
            <div id="eliminar">
                 <button id="eliminar-boton" onClick={() => deleteComment(comment.id)}> X </button>
            </div>
        </div>
    )

}


async function getCommentsByUserId(id) {
    return fetch(`${URLITEMS}/users/${id}/comments`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        }
    }).then(response => response.json())
}



async function setComments(setUserComments, userId) {
    await getCommentsByUserId(userId).then(response => {
        response != null ? setUserComments(response) : setUserComments([]);
    })
}



function MyComments() {
    const [user, setUser] = useState({});
    const [isLogged, setIsLogged] = useState(false);
    const [userComments, setUserComments] = useState([])

    if (Cookie.get("user_id") > -1 && !isLogged) {
        getUserById(Cookie.get("user_id")).then(response => setUser(response))
        setIsLogged(true)

    }


    if (userComments.length <= 0 && Cookie.get("user_id") > -1) {
        setComments(setUserComments, Cookie.get("user_id"))
    }

    const error = (
        <div>
            <div> BOO ERROR :(((( </div>
            <div> There's no comments yet :D </div>
        </div>
    )

    const logreg = (
        <div>
            <a id="login" onClick={() => goto("/login")}>Login</a>
            <a id="register" onClick={() => goto("/register")}>Register</a>
        </div>
    )

    const loggedout = (
        <div>
            <a id="logout" onClick={logout}> <span> Welcome in {user.first_name} </span> </a>
        </div>
    )

    return (
        <div className="comments">
            <div className="topnavHOME">
                <img src={logo} width="80px" height="80px" id="logo" onClick={() => goto("/")} />
            </div>

            <div id="mySidenav" className="sidenav" >
                {isLogged ? loggedout : logreg}
                <a id="sistema" onClick={() => goto("/sistema")}>Sistema</a>
                <a id="publications" onClick={() => goto("/publications")}>Publicaciones</a>
                <a id="mycomments" className="clicked" onClick={() => goto("/mycomments")}>Mis Comentarios</a>
            </div>

            <div id="main">
                {userComments.length > 0 ? showComments(userComments) : error}
            </div>
        </div>
    );
}

export default MyComments;