<script lang="ts">
	import "../../app.css";
	import { goto } from "$app/navigation";

	let password: string = "";
	let email: string = "";
	let name: string = "";
	let surname: string = "";
	let tg: string = "";

	function validateForm() {
		return (
			email.length >= 6 &&
			password.length >= 6 &&
			name.trim() !== "" &&
			surname.trim() !== "" &&
			tg.trim() !== ""
		);
	}

	let id: number;
	import { api } from "../api.js";
	async function create_user() {
		if (!validateForm()) {
			alert("Все поля должны быть заполнены! Пароль должен содержать больше 6 символов.");
			return;
		} else {
			let response = await fetch(api + "/user/create", {
				method: "POST",
				body: JSON.stringify({ PWD: password, email: email, name: name, surname: surname, tg: tg }),
				headers: {
					"Content-Type": "application/json"
				}
			});
			if (!response.ok) {
				const errorObj = await response.json();
				if (response.status === 500) {
					alert("Произошла ошибка. Возможно, пользователь с таким email уже существует.");
				}
			}

			let obj = await response.json();

			id = obj.id;
			handleSubmit();
		}
	}
	
	let selectedValue = "yes";
	async function handleSubmit() {
		if (selectedValue === "yes") {
			goto(`/form/${id}`);
		} else {
			goto(`/main/${id}`);
		}
	}

</script>

<header>
	<a href="/">
		<img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
	</a>
</header>
<h1>Регистрация</h1>
<main>
	<form>
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
		<div class="question">
			<div><label for="name">Имя</label></div>
			<div><input bind:value={name} placeholder="Иван" id="name" type="text" /></div>
		</div>
		<div class="question">
			<div><label for="surname">Фамилия</label></div>
			<div><input bind:value={surname} placeholder="Иванов" id="surname" type="text" /></div>
		</div>
		<div class="question">
			<div><label for="tg">Телеграмм</label></div>
			<div><input bind:value={tg} placeholder="@meow" id="tg" type="text" /></div>
		</div>
		<fieldset>
			<legend>Создать анкету:</legend>
			<label for="yes">да</label>
			<input
				class="choice"
				id="yes"
				type="radio"
				name="val"
				value="yes"
				bind:group={selectedValue}
			/>
			<label for="no">нет</label>
			<input class="choice" id="no" type="radio" name="val" value="no" bind:group={selectedValue} />
		</fieldset>

		<div class="submit">
			<input type="submit" value="Зарегистрироваться" on:click={create_user} id="submit" />
		</div>
	</form>
</main>

<style lang="scss">
	h1 {
		display: flex;
		justify-content: center;
		margin: 12px;
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
			width: 578px;
			height: 572px;
			border-radius: 8px;
			margin: 15px;

			input {
				background-color: #1e1e1e;
				width: 530px;
				border-radius: 8px;
				color: aliceblue;
				height: 40px;
			}
			fieldset {
				display: flex;
				justify-content: center;
				margin: 10px;
				width: 300px;
				height: 60px;
				.choice {
					margin: 10px;
					height: 10px;
					width: auto;
				}
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
			#submit:disabled {
				background-color: gray;
				color: lightgray;
				cursor: not-allowed;
			}
		}
	}
</style>
