<script lang="ts">
	import "../../../app.css";
	import "../../api";
	export let data;
	let id = data.user.ID;

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
	import { api } from "../../api.js";

	async function list_forms() {
		let response = await fetch(api + "/form/all", {
			headers: {
				"Content-Type": "application/json"
			}
		});
		let obj = await response.json();
		forms = obj.forms;
		getFormBySphere();
	}

	const getFormBySphere = () => {
		if (selectedSphere === "all") {
			filteredSpeheres = forms;
		} else {
			filteredSpeheres = forms.filter(summary => summary.sphere === selectedSphere); // Фильтрация по сфере
		}
	};

	onMount(() => {
		list_forms();
	});
	let filters = {
		designer: false,
		frontend: false,
		backend: false,
		ml: false
	};
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

	// Слежение за изменением selectedSphere, чтобы применить фильтрацию
	$: selectedSphere, getFormBySphere();

	// Загружаем формы при монтировании компонента
	onMount(() => {
		list_forms();
	});
</script>

<header>
	<a href="/">
		<img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
	</a>
	<div>
		<a href="/form/{id}">Создать анкету</a>
		<a href="#job_market">Рынок вакансий</a>

		<a href="/create_team/{id}">Создать команду</a>
	</div>
</header>
<main>
	<div class="logo">
		<h1>Dream team</h1>
		<h2>найди свою команду мечты</h2>
	</div>

	<div class="subheading">
		<h3>Рынок вакансий</h3>

		<form>
			<select id="sphere-select" bind:value={selectedSphere}>
				<option value="all"> Все </option>
				<option value="designer">Дизайнер</option>
				<option value="frontend">Фронтенд</option>
				<option value="ml">ML</option>
				<option value="backend">Бэкенд</option>
			</select>
		</form>
	</div>
	<div class="job_market" id="job_market">
		<div>
			<ul class="questionnaires">
				{#each filteredSpeheres as { ID, name, photo, about, sphere }}
					<li class="questionnaire">
						<img src={photo} alt="" width="384px" height="400px" />
						<h3>{name}, {sphere}</h3>
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

				border-radius: 8px;
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

			#sphere-select {
				color: aliceblue;
				width: 243px;
				height: 37px;
				border-radius: 10px;
				background-color: rgba(44, 44, 44, 1);
			}
		}
		.job_market {
			width: 100%;
			min-height: 1000px;
			display: flex;
			flex-direction: row;
			background-color: #1e1e1e;

			.questionnaires {
				margin-left: 13px;
				display: flex;
				flex-direction: row;
				flex-wrap: wrap;
				justify-content: center;

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
	@media (max-width: 578px) {
		header {
			height: auto;
			img {
				height: 15px;
				margin-top: 7px;
			}
			div {
				margin-top: 13px;
				flex-wrap: wrap;
				justify-content: end;
			}
		}
		main {
			.logo {
				h1 {
					text-align: center;
					margin: 0;
					font-size: 62px;
				}
				h2 {
					text-align: center;
					margin: 0;
					font-size: 20px;
				}
			}
			.subheading {
				flex-direction: column;
				align-items: center;
			}
			.job_market {
				flex-direction: column;
				align-items: center;
				justify-content: center;
				.questionnaires {
					flex-direction: column;
					align-items: center;
					justify-content: center;
				}
			}
		}
		@media (max-width: 348px) {
			main {
				width: 100%;
				.subheading {
					#sphere-select {
						width: 170px;
					}
				}
				.job_market {
					flex-direction: column;
					width: 100%;
					margin: 0;

					.questionnaires {
						flex-direction: column;
						align-items: center;
						width: 100%;
						margin: 0;
						gap: 0;
						.questionnaire {
							width: 100%;
							margin: 5px;
						}
					}
				}
			}
		}
	}
</style>
