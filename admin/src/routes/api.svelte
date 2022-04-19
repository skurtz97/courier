<script lang="ts">
  import RouteList from "../components/route-list.svelte";
  import FetchInput from "../components/fetch-input.svelte";
  import FetchOutput from "../components/fetch-output.svelte";
  import { routes } from "../stores/route-store";
  import type { Route } from "../types/route";

  let selected: Route = $routes[0];
  let selectedId = 0;

  function setSelected(id: number) {
    selected = $routes[id];
    selectedId = id;
  }
</script>

<main>
  <div class="routes">
    <RouteList routes={$routes} {setSelected} {selectedId} />
    <FetchInput route={selected} />
  </div>
  <div class="output">
    <FetchOutput />
  </div>
</main>

<style>
  main {
    display: flex;
    width: 100%;
  }
  .routes {
    width: 50%;
    height: calc(100% - 1rem);
    display: flex;
    flex-direction: column;
    margin: 0.5rem;
  }
  .output {
    width: 50%;
    margin: 0.5rem;
  }
</style>
