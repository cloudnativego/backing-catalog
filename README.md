[![wercker status](https://app.wercker.com/status/fea9dadf1fd38ffc62d5f4f84489a730/m "wercker status")](https://app.wercker.com/project/bykey/fea9dadf1fd38ffc62d5f4f84489a730)

# Backing Catalog Service

Fake catalog service for the Backing Services chapter.

To build and run:

* Do a `glide install` to fetch all of the dependencies from the `glide.lock` file and put them in the `vendor/` directory.
* `go build` to build the executable honoring the new vendor directory.
* Run the application. Make sure that the fulfillment service is also running, otherwise requests for individual SKU details will fail.

# Service API

| Resource | Method | Description |
|---|---|---|
| /catalog | GET | Retrieves a summary of catalog items |
| /catalog/{sku} | GET | Retrieves details for an individual catalog item. This will invoke the fulfillment backing service |
