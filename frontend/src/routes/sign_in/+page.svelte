<script lang="ts">
	import "../../app.css";
	import { goto } from "$app/navigation";

	let password: string = "";
	let email: string = "";

	function validateForm() {
		return email.length >= 6 && password.length >= 6;
	}

	let id: number;
	import { api } from "../api.js";
	async function login_user() {
		if (!validateForm()) {
			alert("Пожалуйста, заполните все поля правильно.");
			return;
		} else {
			try {
				let response = await fetch(api + "/user/login", {
					method: "POST",
					body: JSON.stringify({ PWD: password, email: email }),
					headers: {
						"Content-Type": "application/json"
					}
				});
				if (!response.ok) {
					throw new Error(`HTTP error! status: ${response.status}`);
				}

				let obj = await response.json();
				id = obj.id;
				goto(`/main/${id}`);
			} catch (error) {
				console.error("Ошибка при запросе:", error);
			}
		}
	}


	//  смена цвета для кнопки
	let color1 = "#ffffff";
	let font_color1 = "black";

	function handleMouseOver(e) {
		color1 = "#1e1e1e";
		font_color1 = "#ffffff";
	}
	function handleMouseOut(e) {
		color1 = "#ffffff";
		font_color1 = "black"
	}
</script>

<header>
	<a href="/">
		<img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
	</a>
</header>
<h1>Вход</h1>
<main>
	<form action="">
		<div class="question">
			<div><label for="email">Почта</label></div>
			<div><input bind:value={email} placeholder="1234@mail.ru" id="email" type="email" /></div>
		</div>
		<div class="question">
			<div><label for="password">Пароль</label></div>
			<div>
				<input bind:value={password} placeholder="Введите пароль" id="password" type="password" />
			</div>
		</div>

		<div class="submit">
			<!-- <input type="submit" on:click={login_user} value="Войти" id="submit" /> -->
			<button
				on:click={login_user}
				id="submit"
				style="background-color: {color1}; color: {font_color1}"
				on:mouseover={handleMouseOver}
				on:mouseout={handleMouseOut}
			>
			Войти</button
			>
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
			// width: 578px;
			width: 41%;
			height: 322px;

			input {
				background-color: #1e1e1e;
				width: 100%;
				border-radius: 8px;
				color: aliceblue;
				height: 40px;
			}

			.question {
				margin: 10px;
				width: 90%;
			}

			button {
				margin-top: 5px;
				background-color: rgba(245, 245, 245, 1);
				height: 190%;
				width: 310px;
				border-radius: 8px;
				color: black;
			}
			// #submit {
			// 	margin-top: 5px;
			// 	height: 190%;
			// 	width: 110%;
			// 	border-radius: 8px;
			// }
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
			button{
			width: 110px;
		}
		}}
		
	}
</style>
