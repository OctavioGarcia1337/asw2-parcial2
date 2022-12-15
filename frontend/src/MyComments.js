import React, {useEffect, useState} from "react";
import "./css/Orders.css";
import logo from "./images/logo.svg"
import Cookies from "universal-cookie";
import "./css/Home.css";
import { HOST, ITEMSPORT, USERSPORT, MESSAGESPORT} from "./config/config";
import Comment from "./Comment"

const Cookie = new Cookies();
const URLITEMS = `${HOST}:${ITEMSPORT}`
const URLUSERS = `${HOST}:${USERSPORT}`
const URLMESSAGES = `${HOST}:${MESSAGESPORT}`

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


async function deleteComment(id) {
    return await fetch(`${URLMESSAGES}/messages/${id}`, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json"
        }
    }).then(response => {
        response.status === 200 ? goto("/mycomments") : alert("error deleting comment");
    })
}

async function getCommentsByUserId(id) {
    return fetch(`${URLMESSAGES}/users/${id}/messages`, {
        method: "GET",
        headers: {
            "Content-Type": "comment/json"
        }
    }).then(response => response.json())
}


function goto(path) {
    window.location = window.location.origin + path
}


function showComments(comments) {

    return comments.map((comment) =>
        <div>
            <div>
                <Comment
                    key={comment.message_id}
                    comment={comment}
                    first_name={comment.first_name}
                />
            </div> 
            <div id="eliminar">
                {!comment.system ? <button id="eliminar-boton" onClick={() => deleteComment(comment.message_id)}> X </button> : void(0)}
            </div>
        </div>
    )

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
    const [needComments, setNeedComments] = useState(true)

    if (Cookie.get("user_id") > -1 && !isLogged) {
        getUserById(Cookie.get("user_id")).then(response => setUser(response))
        setIsLogged(true)

    }

    useEffect(() => {
        if (userComments.length <= 0 && Cookie.get("user_id") > -1 && needComments) {
            setComments(setUserComments, Cookie.get("user_id"))
        }
        setNeedComments(false);
    }, [userComments.length])


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
                <img src={logo} width="80px" height="80px" id="logo" onClick={()=>goto("/")} /> <p> TuCasa.com </p>
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