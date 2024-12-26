<script lang="ts">
	import "../../../app.css";
	import { goto } from "$app/navigation";

	import { onMount } from "svelte";
	export let data;
</script>

<header>
	<a href="/">
		<img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
	</a>
	<div>
		<a href="/captain_main/{data.team.ID}">Рынок вакансий</a>
		<a href="/main/{data.team.ID_Kap}">Режим участника</a>
	</div>
</header>
<main>
	<div class="logo">
		<h1>{data.team.name}</h1>
		<div class="about">
			<h3>Описание:</h3>
			{data.team.about}
		</div>
	</div>

	
	<div class="job_market" id="job_market">
		<div>
			<ul class="questionnaires">
				{#each data.team.members as { ID, name, surname, email, tg }}
					<li class="questionnaire">
						<h3>{name}, {surname}</h3>
						<p>{email}</p>
						<p>{tg}</p>
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
			min-height: 452px;
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
			.about{
				margin: 13px;
				border: #ffffff;
				border-style: solid;
				padding: 13px;
			}
		}
		.job_market {
			padding-top: 30px;
			min-height: 500px;
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
					height: 218px;
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
					button {
						margin-top: auto;
					}
				}
			}
		}
	}
</style>
