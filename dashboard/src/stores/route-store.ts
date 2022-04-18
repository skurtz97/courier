import { writable } from "svelte/store";
import type { Route } from "../types/route";

export const routes = writable<Route[]>([
  {
    id: 0,
    selected: false,
    method: "GET",
    path: "/api/v1/ping",
    description: "Ping the server",
  },
  {
    id: 1,
    selected: false,
    method: "GET",
    path: "/api/v1/posts",
    description: "List all posts",
  },
  {
    id: 2,
    selected: false,
    method: "POST",
    path: "/api/v1/posts",
    description: "Create a new post",
  },
  {
    id: 3,
    selected: false,
    method: "GET",
    path: "/api/v1/posts/:id",
    description: "Get a post by id",
  },
  {
    id: 4,
    selected: false,
    method: "PUT",
    path: "/api/v1/posts/:id",
    description: "Update a post",
  },
  {
    id: 5,
    selected: false,
    method: "DELETE",
    path: "/api/v1/posts/:id",
    description: "Delete a post",
  },
]);
