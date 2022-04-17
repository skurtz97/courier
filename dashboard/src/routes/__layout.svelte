<div class="root">
  <nav>
    <div class="nav-content">
      <ul>
        <li><a href="/">Home</a></li>
        <li><a href="/api">API</a></li>
        <li><a href="/storage">Storage</a></li>
        <li><a href="/logs">Logs</a></li>
        <li><a href="/metrics">Metrics</a></li>
        <li class="spacer">Spacer</li>
        <li class="settings"><a href="/settings">Settings</a></li>
      </ul>
    </div>
  </nav>
  <slot />
</div>

<style>
  :global(body, html, .svelte-body) {
    height: 100%;
    margin: 0;
    padding: 0;
    font-family: "Noto Sans", sans-serif;
    background-color: var(--slate-200);
  }
  :global(:root) {
    --slate-50: #f8fafc;
    --slate-100: #f1f5f9;
    --slate-200: #e2e8f0;
    --slate-300: #cbd5e1;
    --slate-400: #94a3b8;
    --slate-500: #64748b;
    --slate-600: #475569;
    --slate-700: #334155;
    --slate-800: #1e293b;
    --slate-900: #0f172a;

    --sky-50: #f0f9ff;
    --sky-100: #e0f2fe;
    --sky-200: #bae6fd;
    --sky-300: #7dd3fc;
    --sky-400: #38bdf8;
    --sky-500: #0ea5e9;
    --sky-600: #0284c7;
    --sky-700: #0369a1;
    --sky-800: #075985;
    --sky-900: #0c4a6e;

    --red-900: #7f1d1d;
    --red-800: #991b1b;
    --red-600: #dc2626;
    --red-400: #f87171;
    --red-200: #fecaca;

    --green-900: #064e3b;
    --green-700: #047857;
    --green-500: #10b981;
    --green-300: #6ee7b7;
    --green-200: #a7f3d0;

    --orange-700: #b45309;
    --orange-500: #f59e0b;
    --orange-300: #fcd34d;
    --orange-200: #fde68a;
  }

  .root {
    display: flex;
    flex-direction: row;
    height: 100%;
  }

  nav {
    background: var(--slate-600);
    color: var(--sky-100);
    width: 8rem;
  }

  .nav-content {
    font-size: 1.25rem;
    letter-spacing: 0.15rem;
    height: 100%;
    text-align: left;
    padding: 1rem;
  }
  ul {
    display: flex;
    flex-direction: column;
    height: 100%;
    list-style-type: none;
    padding: 0;
  }
  .spacer {
    flex: 1;
    visibility: hidden;
  }
  li {
    padding: 0.5rem 0;
  }

  li > a {
    color: var(--sky-100);
    text-decoration: none;
  }
  li > a:hover {
    color: var(--slate-400);
  }
</style>
