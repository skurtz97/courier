<script lang="ts">
  import Button from "../components/button.svelte";
  import Input from "../components/input.svelte";
  import { InputType } from "../types/input";
  import type { Route, ParamType } from "../types/route";

  export let onSubmit: (e: SubmitEvent) => void;
  export let route: Route;
  function paramTypeToInputType(paramType: ParamType): InputType {
    switch (paramType) {
      case "string":
        return InputType.text;
      case "integer":
        return InputType.number;
      case "boolean":
        return InputType.checkbox;
      default:
        return InputType.text;
    }
  }
</script>

<form class="form" on:submit|preventDefault={onSubmit}>
  <div class="form-heading">
    <span class="form-method">{route.method}</span>
    <span class="form-path">{route.path}</span>
    <span class="form-description">{route.description}</span>
  </div>

  <div class="form-input">
    <h2 class="form-input-heading">URL Parameters</h2>
    {#if route.url_param}
      <Input
        label={`${route.url_param.name}: ${route.url_param.type}`}
        name={route.url_param.name}
        type={paramTypeToInputType(route.url_param.type)} />
    {:else}
      <p class="form-text">No URL parameters</p>
    {/if}
    <h2 class="form-input-heading">Query Parameters</h2>
    {#each route.query_params as param}
      <Input label={param.name} name={param.name} />
    {:else}
      <p class="form-text">No query parameters</p>
    {/each}
  </div>
  <Button>Submit</Button>
</form>

<style>
  .form {
    background-color: var(--slate-300);
    box-shadow: 0px 0px 2px 2px var(--slate-400);
    display: flex;
    flex-direction: column;
    padding: 1rem;
    height: 50%;
  }
  .form-heading {
    font-size: 1.5rem;
    margin-bottom: 1rem;
    padding-bottom: 0.5rem;
    display: flex;
    flex-direction: row;
    border-bottom: 1px solid var(--slate-500);
  }
  .form-method {
    font-weight: bold;
    text-align: start;
    width: 86px;
  }
  .form-path {
    text-align: end;
    flex: 2;
  }
  .form-description {
    text-align: end;
    flex: 2;
  }
  .form-input {
    display: flex;
    flex-direction: column;
    margin-top: 1rem;
    margin-bottom: 1rem;
    flex: 1;
    overflow-y: scroll;
  }
  .form-input-heading {
    font-size: 1.25rem;
    font-weight: bold;
    margin-bottom: 0.5rem;
    margin-top: 1rem;
  }
  .form-input-heading:first-of-type {
    margin-top: 0;
  }
</style>
