<script lang="ts">
	import "../../../app.css";
	import { goto } from "$app/navigation";

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
	import { api } from "../../api";

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

	onMount(() => {
		list_forms();
	});

	let sphere = ["all", "frontend", "backend", "ml", "designer"];
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

	const getFormBySphere = () => {
		if (selectedSphere === "all") {
			filteredSpeheres = forms; // Если выбран фильтр "all", показываем все формы
		} else {
			filteredSpeheres = forms.filter(summary => summary.sphere === selectedSphere); // Фильтрация по сфере
		}
	};
	let filters = {
		designer: false,
		frontend: false,
		backend: false,
		ml: false
	};
	$: selectedSphere, getFormBySphere();

	onMount(() => {
		getFormBySphere();
	});

	export let data;
	let id = data.team.ID;
	let id_user: number;

	async function add_to_team(userId: number) {
		id_user = userId;
		let response = await fetch(
			api + "/team/add/{id_team}/{id_user}?id_team=" + id + "&id_user=" + id_user,
			{
				method: "POST",
				body: JSON.stringify({ id_team: id, id_use: id_user }),
				headers: {
					"Content-Type": "application/json"
				}
			}
		);

		let obj = await response.json();
		console.log(obj);
	}
	// onMount(() => {
	// 	add_to_team(id_user)
	// });
</script>

<header>
	<a href="/">
		<img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
	</a>
	<div>
		<a href="#job_market">Рынок вакансий</a>
		<a href="/team/{id}">Моя команда</a>
		<a href="/main/{data.team.ID_Kap}">Режим участника</a>
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
				{#each filteredSpeheres as { ID, name, photo, about, sphere, ID_User }}
					<li class="questionnaire">
						<img src={photo} alt="" width="384px" height="400px" />
						<h3>{name}, {sphere}</h3>
						<p>{about.substring(0, 500)}</p>
						{#if !data.team.members.find(member => member.ID === ID_User)}
							<button on:click={() => add_to_team(ID_User)}>Пригласить в команду</button>
						{/if}
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
					button {
						margin-top: auto;
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
