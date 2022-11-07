// header
const navToggle = document.querySelector(".nav-toggle")
const navMenu = document.querySelector(".nav-menu")
// search-bar
const itemCardTemplate = document.querySelector("[data-item-template]")
const itemCardContainer = document.querySelector("[data-item-cards-container]")
const searchInput = document.querySelector("[data-search]")
// solr
const SolrNode = ('solr-node');

// solr-connection
var client = new SolrNode({
    host: '127.0.0.1',
    port: '8983',
    core: 'items',
    protocol: 'http'
});


// header-start
navToggle.addEventListener("click", () => {
    navMenu.classList.toggle("nav-menu_visible");

    if (navMenu.classList.contains("nav-menu_visible")){
        navToggle.setAttribute("aria-label","Close menu");
    }else{
        navToggle.setAttribute("aria-label","Open menu");
    }
});
// header-finish

// search-bar-start
let items = []

searchInput.addEventListener("input", e => {
  const value = e.target.value.toLowerCase()
  items.forEach(item => {
    const isVisible =
      item.item_name.toLowerCase().includes(value) ||  //separando parametros de busqueda
      item.description.toLowerCase().includes(value)
    item.element.classList.toggle("hide", !isVisible)
  })
})


// solr-start
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
    console.log(doc); // devuelve la busqueda
  })
}
});
// solr-finish

// ACA HAY QUE CAMBIAR muchas cosas
fetch("https://jsonplaceholder.typicode.com/users") //aca tendriamos que poner la base de datos o el solr, tipo le esta tirando un fetch a eso
  .then(res => res.json())
  .then(data => {
    items = data.map(item => {
      const card = itemCardTemplate.content.cloneNode(true).children[0]
      const header = card.querySelector("[data-header]")
      const image = card.querySelector("[data-image]")
      const body = card.querySelector("[data-body]")
      header.textContent = item.item_name
      image.textContent = item.image
      body.textContent = item.description
      itemCardContainer.append(card)
      return { item_name: item.item_name, description: item.description, element: card }
    })
  })
  // search-bar-finish