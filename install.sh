# install all the gorilla deps just in case
echo "Installing gorilla toolkit"
go get -u github.com/gorilla/context
go get -u github.com/gorilla/mux
go get -u github.com/gorilla/reverse
go get -u github.com/gorilla/rpc
go get -u github.com/gorilla/schema
go get -u github.com/gorilla/securecookie
go get -u github.com/gorilla/sessions
go get -u github.com/gorilla/websocket

# SQL driver. 
echo "Installing SQLx"
go get -u github.com/jmoiron/sqlx
# To simplify middleware and a bit of abstraction
echo "Installing Negroni"
go get -u github.com/codegangsta/negroni
echo "Done! Build and deploy!"