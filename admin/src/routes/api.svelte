<script lang="ts">
	import { onMount } from "svelte";

	import type { Route } from "$types/route";
	import Editor from "$components/editor.svelte";

	// the route we have selected, defaults to the first one
	// this value is set by the index in #each
	let selected = 0;
	let routes: Route[] = [
		{
			method: "GET",
			path: "/api/v1/ping",
			description: "Ping the server",
			url_param: {
				name: "",
				type: "string"
			},
			query_params: null
		}
	];
	onMount(async () => {
		routes = await getRoutes();
	});
	async function getRoutes(): Promise<Route[]> {
		console.log("getting routes");
		const res = await fetch("http://localhost:8080/api/v1/routes");
		const data: Route[] = await res.json();
		console.log(data);
		return data;
	}
</script>

<main>
	<section class="routes">
		<ul>
			{#each routes as route, index}
				<li
					class={`${route.method.toLowerCase()} ${selected === index ? "selected" : ""}`}
					on:click={() => (selected = index)}
				>
					<span class="method">{route.method}</span>
					<span class="path">{route.path}</span>
					<span class="description">{route.description}</span>
				</li>
			{:else}
				<h1>No routes</h1>
			{/each}
		</ul>
		<form>
			<div class="params">
				<div class="input-group">
					<h1>Path Parameters</h1>
					{#if routes && routes[selected].url_param.name !== ""}
						<label for={`pparam_${selected}`}>{routes[selected].url_param.name}</label>
						<input type="text" name={`pparam_${selected}`} />
					{:else}
						<h3>No path parameters</h3>
					{/if}
				</div>
				<div class="input-group">
					<h1>Query Parameters</h1>
					{#if routes && routes[selected].query_params}
						<!--
                        {#each $routes[selected].query_params as param}
                            <label for={`qparam_${selected}`}>{param.name}</label>
                            <input type="text" name={`pparam_${selected}`} />
                        {/each}
                        -->
					{:else}
						<h3>No query parameters</h3>
					{/if}
				</div>
			</div>
			<div class="body">
				<h1>Request Body</h1>
				<Editor height={580} editable="true" />
			</div>
		</form>
	</section>
	<section class="output">
		<Editor height={580} editable="false" />
	</section>
</main>

<style>
	main {
		display: flex;
		flex-direction: row;
		flex: 1;
	}
	section {
		display: flex;
		flex-direction: column;
		flex: 1;
	}
	ul {
		flex: 1;
		padding: 1rem;
		margin: 1rem;
		overflow-y: auto;
		box-shadow: 0px 0px 5px 5px var(--grey);
		background-color: #eef1f1;
	}

	li {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		font-size: larger;
		padding: 0.5rem;
		margin-bottom: 1rem;
		border: 1px solid var(--grey);
		border-radius: 0.25rem;
	}
	li:hover {
		box-shadow: 1px 1px 2px 2px var(--grey);
		background-color: var(--grey);
		cursor: pointer;
	}
	.get {
		border: 1px solid #6bfc6b;
		background-color: #e0fce0;
	}
	.get.selected {
		background-color: #5cf85c;
	}
	.post {
		border: 1px solid #98e0fa;
		background-color: #e2f5fc;
	}
	.post.selected {
		background-color: #6ac2df;
	}
	.put {
		border: 1px solid #f1d247;
		background-color: #fcf3cc;
	}
	.put.selected {
		background-color: #f5d54a;
	}
	.delete {
		border: 1px solid #f08f8f;
		background-color: #fadede;
	}
	.delete.selected {
		background-color: #f17777;
	}
	.method {
		font-weight: bold;
		flex: 1;
	}
	.path {
		flex: 2;
		text-align: end;
	}
	.description {
		flex: 2;
		text-align: end;
	}
	form {
		flex: 1;
		display: flex;
		flex-direction: row;
		padding: 1rem;
		margin: 1rem;
		box-shadow: 0px 0px 5px 5px var(--grey);
		background-color: #eef1f1;
	}
	.params {
		flex: 1;
		display: flex;
		flex-direction: column;
	}
	.input-group {
		flex: 1;
	}
	.body {
		flex: 1;
		display: flex;
		flex-direction: column;
		max-height: 600px;
	}
	.output {
		flex: 1;
		padding: 1rem;
		margin: 1rem;
		box-shadow: 0px 0px 5px 5px var(--grey);
		background-color: #eef1f1;
	}
</style>
