# Backing Catalog

Fake catalog service for the Backing Services chapter.

To build and run:

* Do a `glide install` to fetch all of the dependencies from the `glide.lock` file and put them in the `vendor/` directory.
* `go build` to build the executable honoring the new vendor directory.
* Run the application. Make sure that the fulfillment service is also running, otherwise requests for individual SKU details will fail.
