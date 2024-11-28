<script lang="ts">
	import "../../app.css";
	export let data;
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
	</div>

	<div class="job_market">
		<div class="sidebar" style="width:20%">
			<div class="filter_box" style="margin-left:30px">
				<p>Фильтрация</p>
				<form action="/action_page.php">
					<input type="checkbox" id="designer" name="designer" />
					<label for="designer"> дизайнер</label><br />
					<input type="checkbox" id="frontend" name="frontend" />
					<label for="frontend">фронтенд </label><br />
					<input type="checkbox" id="backend" name="backend" />
					<label for="backend"> бэкенд</label><br /><br />
					<input type="submit" value="Фильтр" />
				</form>
			</div>
		</div>

		<div>
			<input type="search" id="inputSearch" placeholder="Поиск..." title="Поиск по ключевым словам"/>
			<ul class="questionnaires">
				{#each data.summaries as { id, name, image, description }}
					<li class="questionnaire">
						<img src={image} alt="" width="384px" height="400px" />
						<a href="/project/{id}">{name}</a>
						<p>{description.substring(0, 500)}</p>
					</li>
				{/each}
			</ul>
		</div>
	</div>
</main>

<style lang="scss">
	%color_shared {
		background-color: #1e1e1e;
		color: #ffffff;
	}
	header {
		@extend %color_shared;
		height: 96px;
		display: flex;
		align-items: center;
	}
	main {
		@extend %color_shared;
		.logo {
			height: 452px;
			background-color: #2c2c2c;
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
			height: 189px;
			// надо убрать следующее
			// background-color: #383131;
			border-style: solid;
			border-color: rgb(241, 236, 236);
			h3 {
			}
		}
		.job_market {
			height: 2148px;
			display: flex;
			flex-direction: row;

			.sidebar {
				display: flex;
				justify-content: center;
				// align-items: center;
				.filter_box {
					display: flex;
					flex-direction: column;
					width: 240px;
					height: 278px;
					border-style: solid;
					border-color: rgb(241, 236, 236);
					margin-top: 40px;
					p {
						margin-top: 20px;
						margin-left: 10px;
					}
				}
			}
            #inputSearch{
                @extend %color_shared;
                margin: 37px;
                height: 40px;
                width: 347px;
                border-radius: 30px;
            }
			.questionnaires {
				display: flex;
				flex-direction: row;
				flex-wrap: wrap;
				justify-content: start;
				// margin-left: 40px;
				// margin-right:40px;
				margin: 37px;
				gap: 36px;
				// align-items: flex-start;
				// flex-grow: 1;
				.questionnaire {
					display: flex;
					flex-direction: column;
					// flex-wrap: wrap;
					width: 340px;
					border-style: solid;
					border-color: rgb(241, 236, 236);
					padding: 10px;
					align-items: center;
					img {
						padding: 30px;

					}
				}
			}
		}
	}
</style>
