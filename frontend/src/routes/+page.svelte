<script lang="ts">
	import "../app.css";
	export let data;

	let sphere = ["all", "frontend", "backend", "ml", "design"];
	let filteredSpeheres: Array<{ id: number; name: string; image: string; sphere: string; description: string }> = [];
	let selectedSphere = "all";
	import { onMount } from "svelte";

	$: if (selectedSphere) getBooksByLang();
	const getBooksByLang = () => {
		
		if (selectedSphere === "all") {
			return filteredSpeheres = data.summaries;
		} 
		return filteredSpeheres = data.summaries.filter(summary => summary.sphere === selectedSphere);
	}	
	onMount(() => {
		getBooksByLang();
	});


	
</script>

<header>
	<img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
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
				<option value="all"> Все </option>
				<option value="design">Дизайнер</option>
				<option value="frontend">Фронтенд</option>
				<option value="ml">ML</option>
				<option value="backend">Бэкенд</option>
			</select>
		</form>
	</div>
	<div class="job_market">
		<div>
			<ul class="questionnaires">
				{#each filteredSpeheres as { id, name, image, sphere, description }}
					<li class="questionnaire">
						<img src={image} alt="" width="384px" height="400px" />
						<a href="/{id}">{name}, {sphere}</a>
						<p>{description.substring(0, 500)}</p>
					</li>
				{/each}
			</ul>
		</div>
	</div>
</main>

<style lang="scss">
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
			height: 2148px;
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
						color: aliceblue;
					}
				}
			}
		}
	}
</style>
