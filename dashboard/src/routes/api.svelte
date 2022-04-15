<script>
    import RouteTable from '../components/RouteTable.svelte';
</script>

<main>
    <section>
        <div class="route-list">
            <RouteTable />
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
	h1 {
		margin: 0rem;
		font-size: 32px;
		font-weight: bold;
        height: 60px;
        max-height: 60px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: flex-start; 
        padding: 1rem 1rem 1rem 1rem;
		
	}
	
</style>
