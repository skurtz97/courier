<script>
	export let routes = [
		{ method: 'GET', path: '/api/v1/ping', description: 'Ping the server' },
		{ method: 'GET', path: '/api/v1/post', description: 'List many posts' },
		{ method: 'POST', path: '/api/v1/post', description: 'Create a post' },
		{ method: 'GET', path: '/api/v1/post/{id}', description: 'Get a post' },
		{ method: 'PUT', path: '/api/v1/post{id}', description: 'Update a post' },
		{ method: 'DELETE', path: '/api/v1/post/{id}', description: 'Delete a post' }
	];
</script>

<div class="routes">
	<table>
		<thead>
			<tr>
				<th class="method-header">Method</th>
				<th class="endpoint-header">Endpoint</th>
				<th class="description-header">Description</th>
			</tr>
		</thead>
		<tbody>
			{#each routes as route}
				<tr>
					{#if route.method === 'GET'}
						<td style="color:green">{route.method}</td>
					{:else if route.method === 'POST'}
						<td style="color:blue">{route.method}</td>
					{:else if route.method === 'PUT'}
						<td style="color:orange">{route.method}</td>
					{:else if route.method === 'DELETE'}
						<td style="color:red">{route.method}</td>
					{:else}
						<td style="color:black">{route.method}</td>
					{/if}
					<td>{route.path}</td>
					<td>{route.description}</td>
				</tr><tr />{/each}
		</tbody>
	</table>
</div>

<style>
	.routes {
		flex: 19;
		width: calc(100% - 2rem);
		display: flex;
		flex-direction: column;
		padding: 1rem;
	}
	table {
		border-collapse: collapse;
	}
	th {
		text-align: left;
		font-size: 24px;
		border-bottom: 1px solid #222322;
		padding-bottom: 1rem;
	}
	td {
		padding: 1rem 0rem 1rem 0rem;
		font-size: 20px;
	}
	.method-header {
		width: 25%;
	}
	.endpoint-header {
		width: 45%;
	}
	.description-header {
		width: 30%;
	}
</style>
