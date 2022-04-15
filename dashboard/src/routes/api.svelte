<script>
    import Table from '../components/table.svelte';
</script>

<main>
    <section>
        <div class="route-list">
            <Table />
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
    .route-list{
        flex: 5;
    }
    .output {
        flex: 5;
        background-color: #e5eef5   
    }
	
</style>
