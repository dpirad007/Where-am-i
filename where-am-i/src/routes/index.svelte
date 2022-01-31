<script>
	import Table from "../components/Table.svelte";
	import Navbar from "../components/Navbar.svelte";

	import "../global.css";

	const data = [
		{ rollno: 1, name: "Ved Patil", country: "India" },
		{ rollno: 2, name: "Gaurang Pawar", country: "India" },
		{ rollno: 3, name: "Atharva Pathade", country: "India" },
		{ rollno: 4, name: "Dion Pinto", country: "India" },
	];
</script>

<main>
	<Navbar />
	<Table {data} />
</main>

<style>
</style>
