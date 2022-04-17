<script lang="ts">
  import Route from "../components/route.svelte";
  import type { RouteProps } from "src/types/route";
  let routes: RouteProps[] = [
    {
      method: "GET",
      path: "/api/v1/ping",
      description: "Ping the server",
    },
    {
      method: "GET",
      path: "/api/v1/posts",
      description: "List all posts",
    },
    {
      method: "POST",
      path: "/api/v1/posts",
      description: "Create a new post",
    },
    {
      method: "GET",
      path: "/api/v1/posts/:id",
      description: "Get a post by id",
    },
    {
      method: "PUT",
      path: "/api/v1/posts/:id",
      description: "Update a post",
    },
    {
      method: "DELETE",
      path: "/api/v1/posts/:id",
      description: "Delete a post",
    },
  ];
</script>

<ul>
  {#each routes as route}
    <Route {route} />
  {/each}
</ul>

<style>
  ul {
    list-style: none;
    padding: 0rem;
    margin: 1rem 2rem 1rem 0.5rem;
  }
</style>
