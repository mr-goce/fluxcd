# fluxcd

## Overview

This Go package simulates user interactions with an e-commerce website and includes a mock server for testing. The simulation performs various actions such as browsing products, adding items to the cart, viewing the cart, emptying the cart, checking out, and logging out.

## Usage

### Running the mock server
`go run server.go`
The mock server will start on port `8080` and respond to all requests with a simple message.

### Running the simulation
`go run main.go`

The script will perform the following actions sequentially:

 - Make a GET request to the homepage (/).
 - Set a random currency.
 - Browse a random product.
 - View the cart.
 - Add a random product to the cart.
 - Empty the cart.
 - Checkout with generated user data.


### Functions
`getRequest(path string)`
Makes a GET request to the specified path.

`postRequest(path string, data url.Values)`
Makes a POST request to the specified path with the provided form data.

`browseProduct()`
Browses a random product from the predefined list.

`viewCart()`
Views the current items in the cart.

`addToCart()`
Adds a random product with a random quantity to the cart.

`emptyCart()`
Empties the cart.

`checkout()`
Performs the checkout process with generated user and credit card data.

`logout()`
Logs out the user.

`main()`
The main function that seeds the random number generator, performs a sequence of actions simulating user interactions with the e-commerce website.

## Continuous Deployment with Flux
Flux is used for continuous deployment, managing the deployment of applications within a MicroK8s Kubernetes cluster. Follow these steps to set up and configure Flux for your project.
### Prerequisites
  - MicroK8s installed and configured.
  - Flux CLI installed.