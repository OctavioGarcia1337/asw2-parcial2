export const getComments = async (type, id) => { // cambiar por un GET a la BD
  return await fetch(URL + "/comments/"+ type +"/"+ id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
  
  
  return [
    {
      id: "1",
      body: "First comment",
      username: "Jack",
      userId: "1",
      parentId: null,
      createdAt: "2021-08-16T23:00:33.010+02:00",
    },
    {
      id: "2",
      body: "Second comment",
      username: "John",
      userId: "2",
      parentId: null,
      createdAt: "2021-08-16T23:00:33.010+02:00",
    },
    {
      id: "3",
      body: "First comment first child",
      username: "John",
      userId: "2",
      parentId: "1",
      createdAt: "2021-08-16T23:00:33.010+02:00",
    },
    {
      itemid: "4",
      body: "Second comment second child",
      username: "John",
      userId: "2",
      parentId: "2",
      createdAt: "2021-08-16T23:00:33.010+02:00",
    },
  ];
};

export const createComment = async (text, uname, parentId = null) => { //cambiar por un POST
  await fetch(URL + "/comment", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body:{
      body: text,
      parentId: parentId,
      username: uname,
      createdAt: new Date().toISOString()}
  })
  return {
    body: text,
    parentId: parentId,
    username: uname,
    createdAt: new Date().toISOString(),
  };
  
};
