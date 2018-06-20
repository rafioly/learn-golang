# learn-golang
## SUPERMARKET VISUALIZATION
Simple webservice made using Go.

## REQUIREMENTS
* Go
* github.com/gorilla/mux

## INSTRUCTIONS
* Go to learn-golang directory
* Execute ```go build```
* Execute ```./learn-golang```
* The app will be run on port 8000


## ENDPOINTS
* *GET* **/** Root
* *GET* **/list_all** to see all items being sold by supermarket
* *POST* **/add_item** to add item to be sold by supermarket with item name under the param 'name' and price under the param 'price'
* *DELETE* **/** remove item to no be no longer sold by supermarket with item name under the param 'name'
