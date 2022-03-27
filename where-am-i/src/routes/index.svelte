<script>

//   import { onMount } from "svelte";

	import Table from "../components/Table.svelte";
	import Navbar from "../components/Navbar.svelte";
	import { csvGenerator } from "../components/csvGenerator";

	import "../global.css";
	import { Button } from "spaper";

	let tableHeader = ["rollno","name","address"];

	const data = [
		{ rollno: 1, name: "Ved Patil", address: "India" },
		{ rollno: 2, name: "Gaurang Pawar", address: "India" },
		{ rollno: 3, name: "Atharva Pathade", address: "India" },
		{ rollno: 4, name: "Dion Pinto", address: "India" },
	];
	

//   let tableData = [];

//   onMount(async () => {
//     const res = await fetch(`https://jsonplaceholder.typicode.com/posts`);
//     tableData = await res.json();
//     console.log(tableData);
//   });

const downloadHandler = () => {
	let tableKeys = Object.keys(data[0]); //extract key names from first Object
	csvGenerator(data, tableKeys, tableHeader, "svelte_csv_demo.csv");
}

</script>

<main>
	<Navbar />
	<Table {data} />
	 <Button type="secondary" on:click={downloadHandler}>Download</Button>
</main>

<style>
</style>
