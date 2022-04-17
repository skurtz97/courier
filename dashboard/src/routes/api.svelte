<script lang="ts">
  import List from "../components/route-list.svelte";
</script>

<main>
  <div class="routes">
    <List />
  </div>

  <div class="output">
    <p>Output</p>
  </div>
</main>

<style>
  main {
    display: flex;
    width: 100%;
  }
  .routes {
    width: 50%;
  }
  .output {
    width: 50%;
  }
</style>
