<button class="button">
  <slot />
</button>

<style>
  .button {
    background-color: var(--slate-600);
    color: var(--slate-100);
    border: none;
    border-radius: 0.5rem;
    padding: 0.5rem 1rem;
    font-size: 1.25rem;
    cursor: pointer;
    transition: all 0.2s ease-in-out;
  }
  .button:hover {
    background-color: var(--slate-700);
  }
</style>
