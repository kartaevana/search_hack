<script lang="ts">
	import "../app.css";
	// export let data;


	let forms: Array<{
		ID: number;
		name: string;
		surname: string;
		email: string;
		tg: string;
		ID_User: number;
		photo: string;
		about: string;
		sphere: string;
	}> = [];
	import { onMount } from "svelte";
	import { api } from "./api.js";

	async function list_forms() {
			let response = await fetch(api + "/form/all", {
				headers: {
					"Content-Type": "application/json"
				}
			});
			let obj = await response.json();
			console.log(obj)
			// forms = obj;
			forms = obj.forms;
			console.log(forms);
	}

	onMount(() => {
		list_forms();
	});


	let sphere = ["all", "frontend", "backend", "ml", "design"];
	let filteredSpeheres: Array<{
		ID: number;
		name: string;
		surname: string;
		email: string;
		tg: string;
		ID_User: number;
		photo: string;
		about: string;
		sphere: string;
	}> = [];
	let selectedSphere = "all";

	$: if (selectedSphere) getFormBySphere();
	const getFormBySphere = () => {
		if (selectedSphere === "all") {
			return (filteredSpeheres = forms);
		}
		return (filteredSpeheres = forms.filter(summary => summary.sphere === selectedSphere));
	};
	onMount(() => {
		getFormBySphere();
	});


</script>

<header>
	<img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
	<!-- <button on:click={}>Создать анкету</button> -->
	<div>
		<a href="/form">Создать анкету</a>
		<a href="#job_market">Рынок вакансий</a>
		<a href="">Мои анкеты</a>
		<a href="/ceate_team">Создать команду</a>
	</div>
</header>
<main>
	<div class="logo">
		<h1>Dream team</h1>
		<h2>найди свою команду мечты</h2>
	</div>

	<div class="subheading">
		<h3>Рынок вакансий</h3>

		<input type="search" id="inputSearch" placeholder="Поиск..." title="Поиск по ключевым словам" />

		<form>
			<select id="sphere-select" bind:value={selectedSphere}>
			<!-- <select id="sphere-select"> -->
				<option value="all"> Все </option>
				<option value="design">Дизайнер</option>
				<option value="frontend">Фронтенд</option>
				<option value="ml">ML</option>
				<option value="backend">Бэкенд</option>
			</select>
		</form>
	</div>
	<div class="job_market" id="job_market">
		<div>
			<ul class="questionnaires">
				{#each forms as { ID, name, photo, about, sphere }}
					<li class="questionnaire">
						<img src={photo} alt="" width="384px" height="400px" />
						<a href="/{ID}">{name}, {sphere}</a>
						<p>{about.substring(0, 500)}</p>
					</li>
				{/each}
			</ul>
		</div>
	</div>
</main>

<style lang="scss">
	header {
		display: flex;
		justify-content: space-between;
		div {
			display: flex;
			flex-direction: row;
			a {
				width: 144px;
				height: 32px;
				background-color: rgba(44, 44, 44, 1);
				margin: 7px;
				display: flex;
				justify-content: center;
				align-items: center;
				color: aliceblue;
				// border-style: solid;
				border-radius: 8px;
				// border-color: rgb(241, 236, 236);
			}
		}
	}
	main {
		.logo {
			height: 452px;
			display: flex;
			justify-content: center;
			align-items: center;
			flex-direction: column;
			h1 {
				font-size: 86px;
			}
			h2 {
				font-size: 32px;
				color: rgba(255, 255, 255, 0.7);
			}
		}
		.subheading {
			height: 160px;
			display: flex;
			flex-direction: row;
			background-color: #1e1e1e;
			justify-content: center;
			align-items: baseline;
			gap: 20px;
			padding-top: 30px;
			h3 {
				font-size: 29px;
			}
			#inputSearch {
				// margin: 37px;
				color: aliceblue;
				height: 40px;
				width: 347px;
				border-radius: 30px;
				background-color: rgba(44, 44, 44, 1);
			}
			#sphere-select {
				color: aliceblue;
				width: 243px;
				height: 37px;
				border-radius: 10px;
				background-color: rgba(44, 44, 44, 1);
			}
		}
		.job_market {
			min-height: 1000px;
			display: flex;
			flex-direction: row;
			background-color: #1e1e1e;

			.questionnaires {
				margin-left: 13px;
				display: flex;
				flex-direction: row;
				flex-wrap: wrap;
				justify-content: start;
				gap: 36px;
				.questionnaire {
					display: flex;
					flex-direction: column;
					height: 418px;
					width: 326px;
					border-style: solid;
					border-color: rgb(241, 236, 236);
					padding: 10px;
					align-items: left;
					img {
						height: 100px;
						width: 100px;
						// padding: 30px;
						// margin: 10px;
						margin-left: 113px;
						border-radius: 70px;
						margin-top: 10px;
						margin-bottom: 5px;
					}
					a {
						margin-top: 7px;
						color: aliceblue;
						font-size: 25px;
					}
				}
			}
		}
	}
</style>
