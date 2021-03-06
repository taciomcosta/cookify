# Cookify

_**A WebAPI with delicious recipes and funny GIFS**_

Cookify lets you search for recipes that have your favorite ingredients.

It adds funny gifs to your recipes, making cooking funnier than ever.

## Running
#### docker-compose
You can run the API with a single docker-compose command:
`docker-compose up --build`

But before that, don't forget to add environment variables
to the compose file in the root folder.

You can look at .env.example to see which variables are expected.

#### make
You can also run the API using `make`. 

There are many other useful `make` commands to help you contribute with
the project, like: `make test`, `make lint`.

## Swagger

You can find the available endpoints at `http://localhost:3000/swagger/`

---

_This project is a challenge proposed by [DeliveryMuch](https://www.deliverymuch.com.br/) :hearts:_
