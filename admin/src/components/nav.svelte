<nav class="nav">
  <ul class="nav__list">
    <li class="nav__item"><a class="nav__link" href="/">Home</a></li>
    <li class="nav__item"><a class="nav__link" href="/api">API</a></li>
    <li class="nav__item"><a class="nav__link" href="/storage">Storage</a></li>
    <li class="nav__item"><a class="nav__link" href="/logs">Logs</a></li>
    <li class="nav__item"><a class="nav__link" href="/metrics">Metrics</a></li>
    <li class="nav__item"><a class="nav__link" href="/settings">Settings</a></li>
  </ul>
</nav>

<style>
  .nav {
    background: var(--slate-600);
    color: var(--sky-100);
    width: 8rem;
    min-width: 8rem;
    display: flex;
    flex-direction: column;
  }
  .nav__list {
    display: flex;
    flex-direction: column;
    list-style-type: none;
    padding: 1rem 0.5rem 0.5rem 1rem;
  }
  .nav__list > li {
    margin: 0rem 0rem 0.5rem 0rem;
  }
  .nav__link {
    color: var(--sky-100);
    text-decoration: none;
    font-size: 1.25rem;
  }
  .nav__link:hover {
    color: var(--slate-400);
  }
</style>
