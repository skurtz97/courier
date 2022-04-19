<script lang="ts">
  import type { Route } from "src/types/route";
  export let route: Route;
  export let setSelected: (id: number) => void;
  export let isSelected: boolean;
</script>

<li
  class={`${route.method.toLowerCase()} ${isSelected ? `${route.method.toLowerCase()}-selected` : ""}`}
  on:click={() => setSelected(route.id)}>
  <div class="text">
    <span class="method">{route.method}</span>
    <span class="path">{route.path}</span>
    <span class="desc">{route.description}</span>
  </div>
</li>

<style>
  li {
    color: #000;
    padding: 0.5rem;
    margin-bottom: 0.25rem;
    font-size: large;
    box-shadow: 1px 1px 1px 1px rgb(202, 200, 200);
    border-radius: 0.25rem;
    border: 1px solid var(--slate-300);
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }
  li:hover {
    cursor: grab;
  }
  .text {
    display: flex;
    flex-direction: row;
    flex: 1;
  }
  .get {
    background-color: var(--green-200);
  }
  .get:hover {
    background-color: var(--green-300);
  }
  .get-selected {
    background-color: var(--green-300);
  }

  .post {
    background-color: var(--sky-200);
  }
  .post:hover {
    background-color: var(--sky-300);
  }
  .post-selected {
    background-color: var(--sky-300);
  }
  .put {
    background-color: var(--orange-200);
  }
  .put:hover {
    background-color: var(--orange-300);
  }
  .put-selected {
    background-color: var(--orange-300);
  }
  .delete {
    background-color: var(--red-200);
  }
  .delete:hover {
    background-color: var(--red-300);
  }
  .delete-selected {
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
</style>
