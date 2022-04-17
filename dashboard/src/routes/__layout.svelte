<div class="root">
  <nav>
    <div class="nav-content">
      <ul>
        <div class="nav-top">
          <li><a href="/">Home</a></li>
          <li><a href="/api">API</a></li>
          <li><a href="/storage">Storage</a></li>
          <li><a href="/logs">Logs</a></li>
          <li><a href="/metrics">Metrics</a></li>
        </div>
        <div class="nav-bottom">
          <li class="settings"><a href="/settings">Settings</a></li>
        </div>
      </ul>
    </div>
  </nav>
  <slot />
</div>

<style>
  /* CSS Reset */
  :global(html, body, div, span, applet, object, iframe, h1, h2, h3, h4, h5, h6, p, blockquote, pre, a, abbr, acronym, address, big, cite, code, del, dfn, em, img, ins, kbd, q, s, samp, small, strike, strong, sub, sup, tt, var, b, u, i, center, dl, dt, dd, ol, ul, li, fieldset, form, label, legend, table, caption, tbody, tfoot, thead, tr, th, td, article, aside, canvas, details, embed, figure, figcaption, footer, header, hgroup, menu, nav, output, ruby, section, summary, time, mark, audio, video) {
    margin: 0;
    padding: 0;
    border: 0;
    font-size: 100%;
    font: inherit;
    vertical-align: baseline;
  }

  /* Color definitions */
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
    --red-300: #fca5a5;
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

  /* Make sure svelte generated containers are full height, also set global baseline font and bg color */
  :global(body, html, .svelte-body) {
    height: 100%;
    margin: 0;
    padding: 0;
    font-family: "Noto Sans", sans-serif;
    background-color: var(--slate-200);
  }

  /* Same as above, make sure the root element we defined is full height */
  .root {
    display: flex;
    flex-direction: row;
    height: 100%;
  }

  nav {
    background: var(--slate-600);
    color: var(--sky-100);
    width: 8rem;
    height: 100%;
  }

  .nav-content {
    font-size: 1.25rem;
    letter-spacing: 0.15rem;
    height: 100%;
    text-align: left;
  }
  ul {
    display: flex;
    flex-direction: column;
    list-style-type: none;
    justify-content: space-between;
    padding: 1rem;
    height: calc(100% - 2rem);
  }

  li > a {
    color: var(--sky-100);
    text-decoration: none;
  }
  li > a:hover {
    color: var(--slate-400);
  }
</style>
