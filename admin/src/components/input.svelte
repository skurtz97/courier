<script lang="ts">
  import { InputType } from "../types/input";
  export let label: string;
  export let name: string;
  export let type: InputType = InputType.text;
  export let placeholder: string = "";
</script>

<div class="input-container">
  <label class="input-label" {label} for={name}>{label}</label>
  <input class="input" {type} {name} {placeholder} />
</div>

<style>
  .input-container {
    display: flex;
    flex-direction: column;
  }
  .input-label {
    font-size: 1.25rem;
    margin-bottom: 0.5rem;
  }
  .input {
    font-size: 1.25rem;
    border: none;
    border-bottom: 1px solid var(--slate-500);
    padding: 0.5rem;
    width: max-content;
  }
</style>
