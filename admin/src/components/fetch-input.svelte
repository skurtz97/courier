<script lang="ts">
  import Button from "../components/button.svelte";
  import type { Route } from "../types/route";
  export let route: Route;
</script>

<div class="input">
  <div class="route-text">
    <span class="method">{route.method}</span>
    <span class="path">{route.path}</span>
    <span class="description">{route.description}</span>
  </div>

  <form>
    <div class="input-group">
      <h2>URL Parameter</h2>
      {#if route.url_param !== null}
        <div class="labeled-input">
          <label label="url-parameter" for="url-parameter">id: integer</label>
          <input type="number" name="url-parameter" />
        </div>
      {:else}
        <div>
          <h2>No URL parameter</h2>
        </div>
      {/if}
    </div>
    <div class="input-group">
      <h2>Query Parameters</h2>
      {#each route.query_params as param}
        <div class="labeled-input">
          <label label={param.name} for={param.name}>{param.name}: {param.type}</label>
          <input type="text" name={param.name} />
        </div>
      {:else}
        <div>
          <h2>No Query Parameters</h2>
        </div>
      {/each}
    </div>
    <Button>Submit</Button>
  </form>
</div>

<style>
  .input {
    padding: 1rem;
    height: 50%;
    background-color: var(--slate-300);
    box-shadow: 0px 0px 2px 2px var(--slate-400);
    overflow-y: scroll;
    display: flex;
    flex-direction: column;
  }
  .route-text {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    margin-bottom: 2rem;
    padding-bottom: 1rem;
    border-bottom: 1px solid var(--slate-500);
    font-size: 1.5rem;
  }

  .route-text > .method {
    font-weight: bold;
    text-align: start;
    width: 86px;
  }
  .route-text > .path {
    text-align: end;
    flex: 2;
  }
  .route-text > .description {
    text-align: end;
    flex: 2;
  }

  .input-group {
    margin-bottom: 1rem;
  }
  .input-group > h2 {
    font-size: 1.5rem;
    font-weight: bold;
    margin-bottom: 1rem;
  }
  .labeled-input {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    margin-bottom: 1rem;
  }
  .labeled-input > label {
    font-size: 1.25rem;
    margin-bottom: 0.5rem;
  }
  .labeled-input > input {
    font-size: 1.25rem;
    padding: 0.5rem;
    border-radius: 0.5rem;
    background-color: var(--slate-200);
    border: none;
  }
</style>
