<nav>
	<ul>
		<li><a href="/">Home</a></li>
		<li><a href="/api">API</a></li>
		<li><a href="/storage">Storage</a></li>
		<li><a href="/logs">Logs</a></li>
		<li><a href="/metrics">Metrics</a></li>
		<li><a href="/settings">Settings</a></li>
	</ul>
</nav>

<style>
	nav {
		display: flex;
		flex-direction: column;
		width: 9rem;
		min-width: 9rem;
		background: var(--darkblue);
		color: var(--white);
	}
	ul {
		display: flex;
		flex-direction: column;
		list-style-type: none;
		padding: 1rem 0.5rem 0.5rem 1rem;
	}
	li {
		margin-bottom: 0.5rem;
	}
	a {
		text-decoration: none;
		color: var(--white);
		font-size: x-large;
	}
	a:hover {
		color: var(--blue);
	}
</style>
