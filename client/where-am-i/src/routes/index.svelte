<script>
	import Navbar from "../components/Navbar.svelte";

	const data = [
		{ rollno: 1, name: "Ved Patil", country: "India" },
		{ rollno: 2, name: "Gaurang Pawar", country: "India" },
		{ rollno: 3, name: "Atharva Pathade", country: "India" },
		{ rollno: 4, name: "Dion Pinto", country: "India" },
	];
</script>

<main>
	<Navbar {data} />
</main>

<style>
	main {
		position: relative;
		width: 100%;
		height: 100%;
		color: #333;
		margin: 0;
		padding: 12px;
		box-sizing: border-box;
	}
</style>
