# bookstore
micro service in golang

# usage
```
git clone https://github.com/defsky/bookstore.git
cd bookstore/user
make build
cd ../user-cli
make build
cd ..
docker-compose build
docker-compose up -d
docker exec -it bookstore_user-cli_1 /bin/bash
./user-cli
```