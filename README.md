## Quotes Checker

<details>
      <summary>To run this project you will need</summary>

* Golang IDE

* Postgres Container
    * To install container with all needed variables just run these command:  docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_HOST_AUTH_METHOD=trust -d postgres
    * [You can also refer to this guide if needed](https://hub.docker.com/_/postgres)

</details>

## About Project
This project provides several APIs(in all that API-s you can pass any currency from [external API provider](https://www.frankfurter.app/docs/)) and also along with APIs its providing mechanism of automatic updates of quotes
<details>
      <summary>API-s</summary> 

1. PUT: http://{HOST}:{PORT}/v1/quote?currency=USD/EUR
    * This API allows you to request an update for quotes. In response, you will receive an ID, which you can use to retrieve the result of that update.
    * You must specify the currency you want to update in the query parameters; otherwise, you will receive an error message

2. GET API: http://{HOST}:{PORT}/v1/quotes/{id}
    * In this API, you should pass the ID obtained from the first API. The result will be a quote.
    * GET: http://{HOST}:{PORT}/v1/quotes/{id} - in this API you should pass ID, that you can get from first API and as a result you will have a quote
    * GET: http://{HOST}:{PORT}/v1/quote/latest?currency=USD/EUR - in this API we can request a result of last update of a quote, you should also pass currency as a query parameters or you will get an error with error message
3. GET API: http://{HOST}:{PORT}/v1/quote/latest?currency=USD/EUR
    * This API allows you to request the result of the last quote update. Again, you must pass the currency as a query parameter; otherwise, an error message will be returned
</details>

<details>
      <summary>Automatic updates</summary>
      
1. Automatic updates were implemented using cron jobs with a configurable timer for updates. The list of currencies with automatic updates includes: “EUR/USD”, “EUR/MXN”, “USD/EUR”, “USD/MXN”, “MXN/USD”, and “MXN/EUR”. Note that other currencies are still available using the APIs, but for those, automatic updates are not enabled

   
</details>
