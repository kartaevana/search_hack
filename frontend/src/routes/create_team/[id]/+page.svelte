<script lang="ts">
	import "../../../app.css";
	import { goto } from "$app/navigation";
	export let data;

	let name: string = "";
	let description: string = "";

	function validateForm() {
		// const emailPattern = /^[^s@]+@[^s@]+\.[^s@]+$/;
		return description.trim() !== "" && name.trim() !== "";
	}

	import { api } from "../../api.js";
	let id = data.user.ID;
	async function create_team() {
		if (!validateForm()) {
			// event.preventDefault();
			alert("Пожалуйста, заполните все поля правильно."); // Сообщение об ошибке
			return;
		} else {
			try {
				let response = await fetch(api + "/team/create", {
					method: "POST",
					body: JSON.stringify({ ID_Kap: id, about: description, name: name }),
					headers: {
						"Content-Type": "application/json"
					}
				});
				if (!response.ok) {
					throw new Error(`HTTP error! status: ${response.status}`);
				}

				let obj = await response.json();
				console.log(obj);
				goto(`/captain_main/${id}`); //id команды!!
				id = obj.id;
			} catch (error) {
				console.error("Ошибка при запросе:", error);
			}
		}
	}
</script>

<header>
	<img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
</header>
<h1>Создание команды</h1>
<main>
	<form action="">
		<div class="question">
			<div><label for="name">Название команды</label></div>
			<div><input bind:value={name} placeholder="ДримТим" id="name" type="text" /></div>
		</div>
		<!-- bind:value={about} -->
		<div class="question">
			<div><label for="description">О себе</label></div>
			<div>
				<textarea
					id="description"
					bind:value={description}
					cols="113"
					rows="10"
					placeholder="Я студент 2 курса мисис, начинающий фронтендер ищу команду для хакатона..."
				></textarea>
			</div>
		</div>
		<div class="submit">
			<input type="submit" on:click={create_team} value="Создать команду" id="submit" />
		</div>
	</form>
	<!-- <div><a href="">Забыли пароль?</a></div> -->
</main>

<style lang="scss">
	h1 {
		display: flex;
		justify-content: center;
		margin: 30px;
		font-size: 70px;
	}

	main {
		display: flex;
		justify-content: center;

		form {
			display: flex;
			flex-direction: column;
			justify-content: center;
			align-items: center;
			background-color: #1e1e1e;
			width: 978px;
			height: 372px;

			input {
				background-color: #1e1e1e;
				width: 330px;
				border-radius: 8px;
				color: aliceblue;
				height: 40px;
			}
			input,
			textarea {
				background-color: #1e1e1e;
				width: 929px;
				border-radius: 8px;
				color: aliceblue;
				height: 40px;
				margin-top: 10px;
			}
			textarea {
				height: 100px;
			}

			.question {
				margin: 10px;
			}

			#submit {
				background-color: rgba(245, 245, 245, 1);
				height: 190%;
				width: 310px;
				border-radius: 8px;
				color: black;
			}
		}
	}
</style>
