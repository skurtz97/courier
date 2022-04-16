<script lang="ts">
	export let routes = [
		{ method: 'GET', path: '/api/v1/ping', description: 'Ping the server' },
		{ method: 'GET', path: '/api/v1/post', description: 'List many posts' },
		{ method: 'POST', path: '/api/v1/post', description: 'Create a post' },
		{ method: 'GET', path: '/api/v1/post/{id}', description: 'Get a post' },
		{ method: 'PUT', path: '/api/v1/post{id}', description: 'Update a post' },
		{ method: 'DELETE', path: '/api/v1/post/{id}', description: 'Delete a post' }
	];
</script>

<ul>
	{#each routes as route}
		<li>
			{#if route.method === 'GET'}
				<span style="color:green">{route.method}</span>
			{:else if route.method === 'POST'}
				<span style="color:blue">{route.method}</span>
			{:else if route.method === 'PUT'}
				<span style="color:orange">{route.method}</span>
			{:else if route.method === 'DELETE'}
				<span style="color:red">{route.method}</span>
			{:else}
				<span style="color:black">{route.method}</span>
			{/if}
			<span>{route.path}</span>
			<span>{route.description}</span>
		</li>
	{/each}
</ul>

<style>
	ul {
		list-style: none;
		background-color: #fff;
		padding: 2rem;
		box-shadow: 1px 1px 1px 1px rgb(202, 200, 200);
		border-radius: 0.25rem;
	}
	li {
		color: #000;
		padding: 1rem;
		font-size: large;
	}
</style>
