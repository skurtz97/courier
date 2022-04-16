<script lang="ts">
	import List from '../components/list.svelte';
</script>

<main>
	<section>
		<div class="list-card">
			<List />
		</div>

		<div class="output">
			<p>Output</p>
		</div>
	</section>
</main>

<style>
	main {
		display: flex;
		flex-direction: column;
		width: 100%;
	}
	section {
		display: flex;
		flex-direction: row;
		flex: 10;
	}
	.list-card {
		margin: 0 0 0 0;
		padding: 0;
		flex: 5;
		padding: 2rem;
	}
	.output {
		flex: 5;
		padding: 2rem;
		background-color: #f1f5f9;
	}
</style>
