import {HOST, MESSAGESPORT} from "./config/config";


const URL = HOST + ":" + MESSAGESPORT


export const getComments = async (type, id) => { // cambiar por un GET a la BD
  return await fetch(URL + "/messages", {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}
export const createComment = async (text, uid, itemid) => { //cambiar por un POST

  await fetch(URL + "/message", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body:{
      body: text,
      user_id: uid,
      item_id: itemid,
      system: false,
    }
  }).then(response => response.json())
      .then(response => {
    return {
      body: text,
      user_id: uid,
      item_id: itemid,
      createdAt: response.created_at,
    }
  });
  
};
