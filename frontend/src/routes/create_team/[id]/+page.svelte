<script lang="ts">
	import "../../../app.css";
	import { goto } from "$app/navigation";
	export let data;

	let name: string = "";
	let description: string = "";

	function validateForm() {
		return description.trim() !== "" && name.trim() !== "";
	}

	import { api } from "../../api.js";
	let id = data.user.ID;
	async function create_team() {
		if (!validateForm()) {
			alert("Поля нельзя оставлять пустыми!");
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
				// console.log(obj);
				id = obj.id
				goto(`/captain_main/${id}`); //id команды!!
				id = obj.id;
			} catch (error) {
				console.error("Ошибка при запросе:", error);
			}
		}
	}
</script>

<header>
	<a href="/">
		<img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
	</a>
</header>
<h1>Создание команды</h1>
<main>
	<form>
		<div class="question">
			<div><label for="name">Название команды</label></div>
			<div><input bind:value={name} placeholder="ДримТим" id="name" type="text" /></div>
		</div>
		<div class="question">
			<div><label for="description">О команде:</label></div>
			<div>
				<textarea
					id="description"
					bind:value={description}
					cols="113"
					rows="10"
					placeholder="Мы команда заряженных начинающих хакатонщиков..."
				></textarea>
			</div>
		</div>
		<div class="submit">
			<input type="submit" on:click={create_team} value="Создать команду" id="submit" />
		</div>
	</form>
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
			// width: 978px;
			width: 69%;
			height: 372px;

			input {
				background-color: #1e1e1e;
				// width: 330px;
				// width: 100%;
				border-radius: 8px;
				color: aliceblue;
				height: 40px;
			}
			input,
			textarea {
				background-color: #1e1e1e;
				// width: 929px;
				width: 97%;
				border-radius: 8px;
				color: aliceblue;
				height: 40px;
				margin-top: 10px;
			}
			textarea {
				height: 100px;
				width: 97%;
			}

			.question {
				margin: 10px;
				width: 97%;
			}

			#submit {
				background-color: rgba(245, 245, 245, 1);
				height: 190%;
				width: 310px;
				// width: 100%;
				border-radius: 8px;
				color: black;
			}
		}
	}
	@media (max-width: 780px) {
		h1 {
			font-size: 44px;
		}
		
		main {
			width: 100%;
			form {
				width: 100%;
				margin: 0;
				
			}
		}
	}
	@media (max-width: 320px) {
		main{form{
			#submit{
			width: 150px;
			height: 150%;
		}
		}}
		
	}
</style>
