import { writable } from "svelte/store";
import type { Route } from "../types/route";

export const routes = writable<Route[]>([
  {
    id: 0,
    selected: false,
    method: "GET",
    path: "/api/v1/ping",
    description: "Ping the server",
    url_param: null,
    query_params: null,
  },
  {
    id: 1,
    selected: false,
    method: "GET",
    path: "/api/v1/posts",
    description: "List all posts",
    url_param: null,
    query_params: null,
  },
  {
    id: 2,
    selected: false,
    method: "POST",
    path: "/api/v1/posts",
    description: "Create a new post",
    url_param: null,
    query_params: null,
  },
  {
    id: 3,
    selected: false,
    method: "GET",
    path: "/api/v1/posts/:id",
    description: "Get a post by id",
    url_param: {
      name: "id",
      type: "integer",
    },
    query_params: null,
  },
  {
    id: 4,
    selected: false,
    method: "PUT",
    path: "/api/v1/posts/:id",
    description: "Update a post",
    url_param: {
      name: "id",
      type: "integer",
    },
    query_params: null,
  },
  {
    id: 5,
    selected: false,
    method: "DELETE",
    path: "/api/v1/posts/:id",
    description: "Delete a post",
    url_param: {
      name: "id",
      type: "integer",
    },
    query_params: null,
  },
]);
