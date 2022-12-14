# Arquitectura de Computadoras II - Proyecto Final

##  **Descripcion del proyecto:**
   El trabajo pide desarrollar un sistema de publicación de clasificados, mediante el cual las empresas inmobiliarias puedan cargar sus bases de datos con el posteo de un archivo json de la información relacionada a los inmuebles. Los navegantes pueden buscar esos clasificados desde la home del sitio en base a una oración y traiga los resultados priorizados que permitan ver el detalle de la publicación.

## Endpoints

### BÚSQUEDA

**GET** - /search=:searchQuery

Ejemplo:  /search=vendedor_juan

Response:

    {
	    ...
	    "precio_base": 10000,
	    "vendedor":"Juan",
	    "barrio":"Nueva Cordoba",
	    ...
    }
    
***

**GET** - /searchAll=:searchQuery

Ejemplo:  /searchAll=juan

Response:

    {
	    ...
	    "precio_base": 10000,
	    "vendedor":"Juan",
	    "barrio":"Nueva Cordoba",
	    ...
    }
    
***

**GET** - /items/:id

Ejemplo:  /items=7b1227b0-75cc-4793-874f-f17939803ece

Response:

    {
	    "id": "7b1227b0-75cc-4793-874f-f17939803ece",
	    "titulo": "Pozo dpto. Las Venturas A-I",
	    "tipo":"Departamento",
	    ...
    }
    
***

### ITEMS

**GET** - /items/:item_id

Ejemplo:  /items/7b1227b0-75cc-4793-874f-f17939803ece

Response:

    {
	    "id": "7b1227b0-75cc-4793-874f-f17939803ece",
	    "titulo": "Pozo dpto. Las Venturas A-I",
	    "tipo":"Departamento",
	    ...
    }
    
***

**POST** - /item

Ejemplo:  /item

Body:

    {
	"titulo": "Pozo dpto. Las Venturas A-I",
	"tipo": "Departamento",
	"ubicacion": "Cordoba",
	...
    }
    
***

Response:

    {
	    "id": "7b1227b0-75cc-4793-874f-f17939803ece",
	    "titulo": "Pozo dpto. Las Venturas A-I",
	    "tipo":"Departamento",
	    "ubicacion": "Cordoba",
	    ...
    }
     
***

**POST** - /items

Ejemplo: /items

Body:

    [
	{
		"titulo": "Pozo dpto. Las Venturas A-I",
		"tipo":"Departamento",
		...
	},
	...
	{
		"titulo": "Pozo dpto. Las Venturas B-V",
		"tipo":"Departamento",
		...
	}
    ]

Response:

    [
	{
		"titulo": "Pozo dpto. Las Venturas A-I",
		"tipo":"Departamento",
		...
	},
	...
	{
		"titulo": "Pozo dpto. Las Venturas B-V",
		"tipo":"Departamento",
		...
	}
    ]
    
***

### Messages

**GET**- /messages/:id

 Ejemplo: /messages/2 
 
 Response:

    {
        "messages_id": 2,
        ...
    }
    
***

**GET**-/messages 

 Ejemplo: /messages 
 
 Response:

    {
        "message_id":  2,
        "user_id": 3,
        "body": "...",
        ...
    }
    
***

**GET**-/users/:id/messages 

Ejemplo: /users/:id/messages

Response:

    {
        "user_id": 3,
        ...
    }
    
***

**POST**-/message

Ejemplo: /message

Body:

    {
        "message_id":  2,
        "user_id": 3,
        "body": "...",
        ...
    }
    
***
**DELETE**-/messages/:id

Ejemplo: /messages/53

Resposne: OK!

### Users

**GET**-/users/:id

 Ejemplo: /users/2
 
 Response:

    {
    	"user_id": 2,
    	"username": "pedro",
    	...
    }
    
***

**GET**-/users

 Ejemplo: /users
 
 Response:

    {
    	"user_id": 2,
    	"username": "pedro",
    	...
    }
    
***

**POST**-/user

 Ejemplo: /user
 
 Body:

    {
    	"username": "pedro",
    	"password": "1234",
    	...
    }
    
***

**DELETE**-/user/:id
 
 Ejemplo: /user/3
 
 Response:  OK! 

**POST**-/login
Ejemplo: /login
Body:

    {
    	"username": "juan",
    	"password": "1234",
    	...
    }
    
***

### **BÚSQUEDA:**

Se pidio utilizar un motor de búsqueda que permita una indexación y búsqueda de los ítems por sus características (título, descripción, atributos, ciudad, estado, etc), que se nutra mediante notificaciones del servicio de Items y busque la información de ese servicio.

Documentacion especifica del servicio BÚSQUEDA implementado por nuestro grupo:

Para este microservicio se implemento el motor de busqueda con SOLR a traves de la siguiente imagen de docker:
    
    | docker run -d -p 8983:8983 solr solr-precreate items |

El servico tiene implementado un unico http request - GET Query - Que recibe un string y ejecuta una busqueda a traves del motor de busqueda implementado (SOLR) y devuelve dichos items en un archivo .json. 

A su vez el servicio indexa los items en el motor de busqueda a medida que los items se van cargando en la base de datos implementada por el servicio de ITEMs. Esto se logra utilizando una implementacion de ColasMQTT de RabbitMQ a traves de la siguiente imagen de docker:
    
    | docker run -d --hostname my-rabbit -e RABBITMQ_DEFAULT_USER=user -e RABBITMQ_DEFAULT_PASS=password --name some-rabbit -p 5671:5671 -p 5672:5672 -p 8080:15672 rabbitmq:3-management |
    


### **ITEMS:**

ITEMs tiene la tarea de recibir los datos de los items a medida que son listados, guardandolos en una base de datos. Tambien tiene la funcion de devolver dichos datos. Para un mejor rendimiento (tiempo de respuesta a la hora de devolver datos), implementa una cache local que retiene los datos de los ultimos items manipulados. Por ultimo, realiza la carga de datos de manera asincronica con uso de goRutines.
        ITEMs implementa una base de datos no SQL (mongodb), una cache distribuida (memcached) y una cola de mensajes tipo ColasMQTT (RabbitMQ); a través de sus respectivas imagenes en docker.

 Imagenes de docker:

 - Base de datos MONGO: 
            
         | # docker run -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root --name some-mongo -d mongo:5.0 |
    
    
- Cache Distribuida MEMCACHED:
            
         | docker run --name memcached -p 11211:11211 memcached:1.6.16 |
    
    
- ColasMMQT RabbitMQ:
     
          | Detallado en BÚSQUEDA anteriormente |

En nuestra implementacion el servicio contiene 2 metodos: 

 - POST de items:
                Recibe un Json con los items a cargar.
                Si el json pudo ser procesado, devuelve codigo http 201 (created).
                Asincronicamente, a traves de gorutines, carga uno por uno los items primero en la base de datos
                y luego en la cache.
                Si un item es cargado correctamente en la bdd se envia un mensaje a un cola (implementacion RabbitMQ).
                De lo contrario se carga un log con el mensaje de error.
 - GET de item (por id)
                Recibe itemId como string. 
                Busca el Item en cache.
                De no encontrarlo lo busca en la Base de datos y lo carga en cache.
                Devuelve el item como archivo .json.


     
      http://localhost:3001/d/k6/k6-load-testing-results?orgId=1&refresh=5s 
	
     
          | docker-compose run k6 run /scripts/ewoks.js |
	

### **MESSAGES** 
Este servicio es el encargado de gestionar los mensajes de comentarios como su nombre lo indica y mediante sus GETS permite Acceder a los ids de los mensajes, su contenido, sus usuarios y filtrar por el el mismo. tambien posee una funcion POST con la cual por medio de un body podemos cargar mensajes y por ultimo con el metodo DELETE sumado al id del respectivo mensaje es posible eliminarlos.
### **Users** 
Mediante este servicio gestionamos las altas y bajas de los usuarios, a su vez entre sus metodos encontramos mecanismos para hallar los usuarios por id, setear sus credenciales para ingresar al sitio y un metodo para acceder al mismo, entre otros datos relevantes del usuario 
### **WORKER_ITEMS:**
Este Trabajador esta suscripto a una cola de topics, da la cual lee las solicitudes de eliminacion de items y las envia al servicio de items, quien es el encargado de procesarlas
### **WORKER_SOLR:**
Este otro trabajador esta suscripto a una cola de topics de la cual se informa hacerca de las interacciones con solr y las envia a items para ser procesadas adecuadamente.
### **FRONTEND:**
El Frontend debia contener la vista de inicio con el input de búsqueda, el listado de Items, el detalle de la publicación.

En la implementacion, el frontend simplemente se comunica con el servicio BÚSQUEDA a traves del request http - GET Query - que se especifico anteriormente, obtiene la informacion de los items y la muestra, cargando tambien las imagenes correspondientes


    
    
