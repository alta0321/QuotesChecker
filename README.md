
## Quotes Checker

<details>
      <summary>To run this project you will need</summary>

* Golang IDE

* Postgres Container
    * To install container with all needed variables just run these command:  docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_HOST_AUTH_METHOD=trust -d postgres
    * [You can also refer to this guide if needed](https://hub.docker.com/_/postgres)

</details>


<details>
      <summary>About project</summary> 

* API-s (in all that API-s you can pass any currency from [external API provider](https://www.frankfurter.app/docs/))
    * PUT: http://{HOST}:{PORT}/v1/quote?currency=USD/EUR - in this API we can request an update for quotes, in response we will get ID, that ID we can use to get the result of that update. In this request you should specify currency that you would like to update in the query parameters or you will get an error with error message.
    * GET: http://{HOST}:{PORT}/v1/quotes/{id} - in this API you should pass ID, that you can get from first API and as a result you will have a quote
    * GET: http://{HOST}:{PORT}/v1/quote/latest?currency=USD/EUR - in this API we can request a result of last update of a quote, you should also pass currency as a query parameters or you will get an error with error message

* Automatic updates
    * Automatic updates were created using cron jobs with configurable timer for an update, the list of currencies with automatic updates : "EUR/USD", "EUR/MXN", "USD/EUR", "USD/MXN", "MXN/USD", "MXN/EUR"  *Note that other currencies still available using API-s, but for those you just have mechanism that will update them automatically


</details>
