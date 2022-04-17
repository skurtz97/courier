<script lang="ts">
  import ChevronDown from "../icons/chevron-down.svelte";
  import type { RouteProps } from "src/types/route";
  export let route: RouteProps;
</script>

{#if route.method === "GET"}
  <li class="li-get">
    <span class="span-get">{route.method}</span>
    <span>{route.path}</span>
    <span class="description">{route.description}</span>
    <ChevronDown />
  </li>
{:else if route.method === "POST"}
  <li class="li-post">
    <span class="span-post">{route.method}</span>
    <span>{route.path}</span>
    <span class="description">{route.description}</span>
    <ChevronDown />
  </li>
{:else if route.method === "PUT"}
  <li class="li-put">
    <span class="span-put">{route.method}</span>
    <span>{route.path}</span>
    <span class="description">{route.description}</span>
    <ChevronDown />
  </li>
{:else}
  <li class="li-delete">
    <span class="span-delete">{route.method}</span>
    <span>{route.path}</span>
    <span class="description">{route.description}</span>
    <ChevronDown />
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
  }
  .li-get {
    background-color: var(--green-200);
  }
  .li-post {
    background-color: var(--sky-200);
  }
  .li-put {
    background-color: var(--orange-200);
  }
  .li-delete {
    background-color: var(--red-200);
  }
</style>
