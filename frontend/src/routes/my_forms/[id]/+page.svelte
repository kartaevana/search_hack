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
		<a href="/main/{data.user.ID}">Рынок вакансий</a>
	</div>
</header>
<main>
	<div class="subheading">
		<h2>Моя заявки</h2>
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
			min-height: 160px;
			display: flex;
			flex-direction: column;
			background-color: #1e1e1e;
			justify-content: center;
			align-items: center;
			gap: 20px;
			padding-top: 30px;
			h3 {
				font-size: 29px;
			}
			#inputSearch {
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
