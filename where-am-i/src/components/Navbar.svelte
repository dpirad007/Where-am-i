<script>
    import Table from "./Table.svelte";
    import { Button, Navbar } from "spaper";

    export let data = [];
</script>

<main>
    <Navbar>
        <h3 slot="brand">
            <a href="#">Where-am-i ?</a>
        </h3>
        <ul class="inline">
            <li><a href="#">Class</a></li>
            <li><a href="#">Logs</a></li>
            <li>
                <Button>Github</Button>
            </li>
        </ul>
    </Navbar>
    <Table {data} />
</main>
