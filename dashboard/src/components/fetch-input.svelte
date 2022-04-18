<script lang="ts">
  import type { Route } from "../types/route";
  export let route: Route;
</script>

<div>
  <h1>{`${route.method} ${route.path} ${route.description}`}</h1>
  <form>
    <h2>Route Parameters</h2>
  </form>
</div>

<style>
  div {
    padding: 1rem;
    height: 50%;
    background-color: var(--slate-300);
    box-shadow: 0px 0px 2px 2px var(--slate-400);
    overflow-y: scroll;
    flex: 1;
  }

  h1 {
    font-size: 1.5rem;
    margin: 0;
    padding: 0.5rem;
  }

  h2 {
  }
</style>
