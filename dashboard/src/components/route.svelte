<script lang="ts">
  import ChevronDown from "../icons/chevron-down.svelte";
  import type { RouteProps } from "src/types/route";
  export let route: RouteProps;
</script>

{#if route.method === "GET"}
  <li class="get">
    <div class="text">
      <span class="method">{route.method}</span>
      <span class="path">{route.path}</span>
      <span class="desc">{route.description}</span>
    </div>
    <div class="icon-container">
      <ChevronDown />
    </div>
  </li>
{:else if route.method === "POST"}
  <li class="post">
    <div class="text">
      <span class="method">{route.method}</span>
      <span class="path">{route.path}</span>
      <span class="desc">{route.description}</span>
    </div>
    <div class="icon-container">
      <ChevronDown />
    </div>
  </li>
{:else if route.method === "PUT"}
  <li class="put">
    <div class="text">
      <span class="method">{route.method}</span>
      <span class="path">{route.path}</span>
      <span class="desc">{route.description}</span>
    </div>
    <div class="icon-container">
      <ChevronDown />
    </div>
  </li>
{:else}
  <li class="delete">
    <div class="text">
      <span class="method">{route.method}</span>
      <span class="path">{route.path}</span>
      <span class="desc">{route.description}</span>
    </div>
    <div class="icon-container">
      <ChevronDown />
    </div>
  </li>
{/if}

<style>
  li {
    color: #000;
    padding: 1rem;
    margin: 1rem;
    font-size: large;
    box-shadow: 1px 1px 1px 1px rgb(202, 200, 200);
    border-radius: 0.25rem;
    border: 1px solid var(--slate-300);
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }
  .text {
    width: 75%;
    display: flex;
    flex-direction: row;
  }
  .get {
    background-color: var(--green-200);
  }
  .get:hover {
    background-color: var(--green-300);
  }
  .post {
    background-color: var(--sky-200);
  }
  .post:hover {
    background-color: var(--sky-300);
  }
  .put {
    background-color: var(--orange-200);
  }
  .put:hover {
    background-color: var(--orange-300);
  }
  .delete {
    background-color: var(--red-200);
  }
  .delete:hover {
    background-color: var(--red-300);
  }
  .method {
    font-weight: bold;
    width: 6rem;
  }
  .path {
    flex: 1;
    text-align: end;
  }
  .desc {
    flex: 1;
    text-align: end;
  }
  .icon-container {
    height: 2rem;
    width: 2rem;
  }
</style>
