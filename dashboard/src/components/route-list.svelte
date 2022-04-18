<script lang="ts">
  import type { Route } from "src/types/route";
  import RouteListItem from "./route-list-item.svelte";
  export let routes: Route[];
  export let setSelected: (id: number) => void;
</script>

<ul>
  {#each routes as route, index (route.id)}
    <RouteListItem {route} {setSelected} />
  {/each}
</ul>

<style>
  ul {
    padding: 1rem;
    margin-bottom: 1rem;
    list-style: none;
    background-color: var(--slate-300);
    box-shadow: 0px 0px 2px 2px var(--slate-400);
    overflow-y: scroll;
    flex: 1;
  }
</style>
