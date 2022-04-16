<nav>
	<h1>Dashboard</h1>
	<ul>
		<li><a href="/">Home</a></li>
		<li><a href="/api">API</a></li>
		<li><a href="/storage">Storage</a></li>
		<li><a href="/logs">Logs</a></li>
		<li><a href="/metrics">Metrics</a></li>
		<li class="spacer">Spacer</li>
		<li class="settings"><a href="/settings">Settings</a></li>
	</ul>
</nav>
<slot />

<style>
	:global(body) {
		padding: 0;
		margin: 0;
		height: 100%;
		font-family: 'Noto Sans', sans-serif;
	}
	:global(html) {
		padding: 0;
		margin: 0;
		height: 100%;
	}
	:global(.svelte-body) {
		padding: 0;
		margin: 0;
		height: 100%;
		display: flex;
		background-color: #dce3e9;
	}

	nav {
		background: #ffffff;
		color: #334155;
		flex: 20;
		height: 100%;
		width: 240px;
		min-width: 240px;
		max-width: 240px;
		display: flex;
		flex-direction: column;
		text-transform: uppercase;
		box-shadow: 1px 0px 2px 2px #ffffff;
	}
	h1 {
		margin: 0rem;
		font-size: 32px;
		font-weight: bold;
		height: 60px;
		max-height: 60px;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		padding: 1rem 0rem 1rem 0rem;
		flex: 1;
	}
	ul {
		list-style: none;
		padding: 0;
		margin: 0;
		flex: 19;
		display: flex;
		flex-direction: column;
	}
	.spacer {
		flex: 10;
		visibility: hidden;
	}
	li {
		font-size: 20px;
	}
	li > a:hover {
		color: #7a8696;
	}
	li > a {
		color: #334155;
		text-decoration: none;
		display: flex;
		width: calc(100% - 2rem);
		height: calc(100% - 2rem);
		padding: 1rem 1rem 1rem 1rem;
	}
</style>
