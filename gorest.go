package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
    "fmt"
)

type Message struct {
    Body string
}

type Info struct {
    Username string
    Password string
}

func main() {
    handler := rest.ResourceHandler{
      PreRoutingMiddlewares: []rest.Middleware{
        &rest.CorsMiddleware{
          RejectNonCorsRequests: false,
          OriginValidator: func(origin string, request *rest.Request) bool {
            return true //origin == "http://localhost:8081"
          },
          AllowedMethods:                []string{"GET", "POST", "PUT"},
          AllowedHeaders:                []string{"Accept", "Content-Type", "X-Custom-Header"},
          AccessControlAllowCredentials: true,
          AccessControlMaxAge:           3600,
        },
      },
    }
    handler.SetRoutes(
      &rest.Route{"GET", "/get", GetParameter},
      &rest.Route{"POST", "/post", PostParameter},
      &rest.Route{"OPTIONS", "/post", PostParameter},
    )
    http.ListenAndServe(":8080", &handler)
}

var store = map[string]*Info{}

func GetParameter(w rest.ResponseWriter, r *rest.Request) {

    informations := make([]*Info, len(store))
    i := 0
    for _, information := range store {
        informations[i] = information
        i++
    }

println("<<get")
    w.WriteJson(&informations)
}

func PostParameter(w rest.ResponseWriter, r *rest.Request) {
    fmt.Printf("HEADER: %s\n", r.Header)
    information := Info{}
    err := r.DecodeJsonPayload(&information)

    if err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        println(err.Error())
        return
    }
    if information.Username == "" {
        rest.Error(w, "Name required", 400)
        println("Name required")
        return
    }
    if information.Password == "" {
        rest.Error(w, "Password required", 400)
        println("Password required")
        return
    }
    store[information.Username] = &information

   information = Info{Username:"yod",Password:"go"}

println(information.Username)


    w.WriteJson(&information)
}
