package main

import (
    "net/http"
    "github.com/gorilla/mux"
)


type Route struct {
    Name string
    Method string
    Pattern string
    HandlerFunc http.HandlerFunc
}


type Routes []Route

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)


    for _, route := range routes {
        var handler http.Handler

        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}


var routes = Routes {
    Route {
       "AllUsers",
       "GET",
       "/users",
       AllUsers,
    },
    Route {
       "GetUser",
       "GET",
       "/users/{id}",
       GetUser,
    },
    Route {
       "NewUser",
       "POST",
       "/users",
       NewUser,
    },
    Route {
        "UpdateUser",
        "PATCH",
        "/users/{id}",
        UpdateUser,
    },
    Route {
        "DeleteUser",
        "DELETE",
        "/users/{id}",
        DeleteUser,
    },
    Route {
        "CreatePost",
        "POST",
        "/posts",
        CreatePost,
    },
    Route {
        "GetPost",
        "GET",
        "/posts/{id}",
        GetPost,
    },
    Route {
        "GetPosts",
        "GET",
        "/posts",
        GetPosts,
    },
    Route {
        "UpdatePost",
        "PATCH",
        "/posts/{id}",
        UpdatePost,
    },
    Route {
        "DeletePost",
        "DELETE",
        "/posts/{id}",
        DeletePost,
    },
}

