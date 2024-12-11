<script lang="ts">
	import "../../app.css";
	import { goto } from "$app/navigation";


	let password: string = "";
	let email: string = "";
	let name: string = "";
	let surname: string = "";
	let tg: string = "";

	let selectedValue = "yes"; // Значение по умолчанию
	async function handleSubmit() {
		if (selectedValue === "yes") {
			// Переход на страницу заполнения формы
			goto("/form");
			// window.location.href = "/form";
		} else {
			// Переход на главную страницу
			goto("/");
		}
	}

	let id: number;
	import { api } from "../api.js";
	async function create_user() {
		try {
			let response = await fetch(api + "/user/create", {
				method: "POST",
				body: JSON.stringify({ PWD: password, email: email, name: name, surname: surname, tg: tg }),
				headers: {
					"Content-Type": "application/json"
				}
			});
			if (!response.ok) {
				throw new Error(`HTTP error! status: ${response.status}`);
			}

			let obj = await response.json();
			console.log(obj);
			id = obj.id;
			handleSubmit();
		} catch (error) {
			console.error("Ошибка при запросе:", error);
		}
	}
	// export let data;
	// const { form } = superForm(data.form);

	// 	import { superForm } from 'sveltekit-superforms';

	// 	const { form, validate } = superForm({
	//     email: {
	//       value: '',
	//       validations: {
	//         required: true,
	//         email: true,
	//       },
	//     },
	//   });
</script>

<header>
	<img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
</header>
<h1>Регистрация</h1>
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
		<!-- <fieldset>
			<legend>Роль:</legend>
			<label for="captain">капитан</label>
			<input
				class="choice"
				id="captain"
				type="radio"
				name="val"
				value="captain"
			/>
			<label for="notcaptain">нет</label>
			<input class="choice" id="notcaptain" type="radio" name="val" value="notcaptain" />
		</fieldset> -->

		<div class="submit">
			<input type="submit" on:click={create_user} value="Зарегистрироваться" id="submit" />
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
		}
	}
</style>
