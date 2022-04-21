<script lang="ts">
  export let response;
</script>

<div class="output">
  <pre>{`${JSON.stringify(response, null, 2)}`}</pre>
</div>

<style>
  div {
    display: flex;
    flex-direction: column;
    padding: 1rem;
    overflow-y: scroll;
    background-color: var(--slate-300);
    box-shadow: 0px 0px 2px 2px var(--slate-400);
    max-height: 100%;
  }
</style>
