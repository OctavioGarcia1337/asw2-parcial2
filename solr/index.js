const SolrNode = ('solr-node');
//const docs = ('./docs.json');



var client = new SolrNode({
    host: '127.0.0.1',
    port: '8983',
    core: 'books',
    protocol: 'http'
});

// Add
// const data = {
//   data: 'pone tu data a agregar',
// };

// client.update(data, function(err, result) {
//   if (err) {
//     console.log(err);
//     return;
//   }
//   console.log('Response:', result.responseHeader);
// });

//-------------------------------------------------------------------------------------------

// // Add from files or docs
// docs.forEach((person) => {
//   client.update(person, function(err, result) {
//     if (err) {
//       console.log(err);
//       return;
//     }
//     console.log('Response:', result.responseHeader);
//   });
// });

//-------------------------------------------------------------------------------------------

// // Delete
// const stringQuery = 'id:2';    // delete document with id 2
// const deleteAllQuery = '*';    // delete all
// const objectQUery = {id: 'xxxxxxxxxxxxxxxx'};   // Object query
// client.delete(deleteAllQuery, function(err, result) {
//   if (err) {
//     console.log(err);
//     return;
//   }
//   console.log('Response:', result.responseHeader);
// });

//-------------------------------------------------------------------------------------------

// Search
// Create const searchBarQuery ***
const authorQuery = { //example
    author: 'Pollo'
  };
  
  // const genderQuery = {      
  //   gender: 'Male'
  // };
  
  // Build a search query var
  const searchQuery = client.query()
    .q(authorQuery)
    .addParams({
      wt: 'json',
      indent: true
    })
    // .start(1)
    // .rows(1)
  
  client.search(searchQuery, function (err, result) {
    if (err) {
      console.log(err);
      return;
    }
  
    const response = result.response;
    console.log(response);
  
    if (response && response.docs) {
      response.docs.forEach((doc) => {
        console.log(doc);
      })
    }
  });