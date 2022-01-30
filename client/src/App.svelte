<script>
	import { Button, Navbar, Table } from "spaper";

	const data = [
		{ rollno: 1, name: "Ved Patil", country: "India" },
		{ rollno: 2, name: "Gaurang Pawar", country: "India" },
		{ rollno: 3, name: "Atharva Pathade", country: "India" },
		{ rollno: 4, name: "Dion Pinto", country: "India" },
	];
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

<style>
</style>
