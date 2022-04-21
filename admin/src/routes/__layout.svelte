<script lang="ts">
	import "$styles/global.css";
	import Sidebar from "$components/sidebar.svelte";
</script>

<Sidebar />
<slot />

<style>
	/* Clean https://colorffy.com/palettes/LuBdqz0NkVhvhK9Dmo85 */
	:global(html) {
		--white: #f7f7f7;
		--grey: #dde6e8;
		--lightblue: #9ed8ff;
		--blue: #0099ff;
		--darkblue: #222f3d;
	}
	:global(.svelte-body) {
		margin: 0;
		padding: 0;
		font-family: "Noto Sans", sans-serif;
		background: var(--white);
		display: flex;
		flex-direction: row;
		min-height: 100vh;
	}
</style>
