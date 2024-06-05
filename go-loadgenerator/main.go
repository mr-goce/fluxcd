package main

import (
    "fmt"
    "math/rand"
    "time"
    "net/http"
    "net/url"

    "github.com/go-faker/faker/v4"
)
var client = &http.Client{}
var products = []string{
    "0PUK6V6EV0",
    "1YMWWN1N4O",
    "2ZYFJ3GM2N",
    "66VCHSJNUP",
    "6E92ZMYYFZ",
    "9SIQT8TOJO",
    "L9ECAV7KIM",
    "LS4PSXUNUM",
    "OLJCESPC7Z",

}
type CheckoutData struct {
    Email                    string `faker:"email"`
    StreetAddress            string `faker:"street_address"`
    ZipCode                  string `faker:"zipcode"`
    City                     string `faker:"city"`
    State                    string `faker:"state_abbr"`
    Country                  string `faker:"country"`
    CreditCardNumber         string `faker:"cc_number"`
    CreditCardExpirationMonth string
    CreditCardExpirationYear  string
    CreditCardCVV            string
}
func getRequest(path string) {
    resp, err := client.Get("http://localhost:8080" + path)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        defer resp.Body.Close()
        fmt.Println("GET", path, "Status:", resp.Status)
    }
}

func postRequest(path string, data url.Values) {
    resp, err := client.PostForm("http://localhost:8080"+path, data)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        defer resp.Body.Close()
        fmt.Println("POST", path, "Status:", resp.Status)
    }
}

func browseProduct() {
    getRequest("/product/" + products[rand.Intn(len(products))])
}

func viewCart() {
    getRequest("/cart")
}

func addToCart() {
    product := products[rand.Intn(len(products))]
    getRequest("/product/" + product)
    postRequest("/cart", url.Values{
        "product_id": {product},
        "quantity":   {fmt.Sprintf("%d", rand.Intn(10) + 1)},
    })
}

func emptyCart() {
    postRequest("/cart/empty", url.Values{})
}

func checkout() {
    addToCart()
    currentYear := time.Now().Year() + 1
    checkoutData := CheckoutData{
        CreditCardExpirationMonth: fmt.Sprintf("%d", rand.Intn(12) + 1),
        CreditCardExpirationYear:  fmt.Sprintf("%d", rand.Intn(70) + currentYear),
        CreditCardCVV:             fmt.Sprintf("%d", rand.Intn(900) + 100),
    }
    faker.FakeData(&checkoutData)
    postRequest("/cart/checkout", url.Values{
        "email":                        {checkoutData.Email},
        "street_address":               {checkoutData.StreetAddress},
        "zip_code":                     {checkoutData.ZipCode},
        "city":                         {checkoutData.City},
        "state":                        {checkoutData.State},
        "country":                      {checkoutData.Country},
        "credit_card_number":           {checkoutData.CreditCardNumber},
        "credit_card_expiration_month": {checkoutData.CreditCardExpirationMonth},
        "credit_card_expiration_year":  {checkoutData.CreditCardExpirationYear},
        "credit_card_cvv":              {checkoutData.CreditCardCVV},
    })
}

func logout() {
    getRequest("/logout")
}
// func main() {
//     rand.Seed(time.Now().UnixNano())
//     fmt.Println("Random number:", rand.Intn(100))
// 	fmt.Println("time", time.Now())
// }

func main() {
    rand.Seed(time.Now().UnixNano())
    getRequest("/")
    currencies := []string{"EUR", "USD", "JPY", "CAD", "GBP", "TRY"}
    postRequest("/setCurrency", url.Values{"currency_code": {currencies[rand.Intn(len(currencies))]}})
    browseProduct()
    viewCart()
    addToCart()
    emptyCart()
    checkout()
    logout()
}