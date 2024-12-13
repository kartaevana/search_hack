<script lang="ts">
	import "../../../app.css";
	import { goto } from "$app/navigation";
	export let data;

	function validateForm() {
		return about.trim() !== "" && selectedImage.trim() !== "" && sphere.trim() !== "";
	}
	let selectedImage = "";

	const images = [
		{ src: "../cat.jpeg", alt: "Cat 1" },
		{ src: "../cat2.jpeg", alt: "Cat 2" },
		{ src: "../cat3.jpeg", alt: "Cat 3" }
	];

	function handleSelection(image) {
		selectedImage = image;
	}

	let ID_User: number;
	let about: string = "";
	let photo: string = "";
	let sphere: string = "";
	import { api } from "../../api.js";
	async function handleSubmit() {
		goto(`/main/${id}`);
	}
	let id = data.user.ID;

	async function create_form() {
		if (!validateForm()) {
			alert("Поля нельзя оставлять пустыми!");
			return;
		} else {
			let response = await fetch(api + "/form/create", {
				method: "POST",
				body: JSON.stringify({
					ID_User: id,
					about: about,
					photo: selectedImage,
					sphere: sphere
				}),
				headers: {
					"Content-Type": "application/json"
				}
			});
			let obj = await response.json();
			handleSubmit();
		}
	}
</script>

<header>
	<a href="/">
		<img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
	</a>
</header>
<h1>Создание анкеты</h1>
<main>
	<form action="">
		<div class="question">
			<div class="question">
				<div><label for="name">Имя</label></div>
				<div>
					<input bind:value={data.user.name} id="name" type="text" readonly />
				</div>
			</div>
			<div class="question">
				<div><label for="surname">Фамилия</label></div>
				<div>
					<input
						bind:value={data.user.surname}
						id="surname"
						type="text"
						readonly
					/>
				</div>
			</div>
			<div class="idk"><label for="role">Роль</label></div>
			<div class="idk">
				<select id="roles" name="roles" bind:value={sphere}>
					<option value="frontend">фронтенд</option>
					<option value="backend">бэкенд</option>
					<option value="designer">дизайнер</option>
					<option value="ml">ML</option>
				</select>
			</div>
		</div>
		<div class="question">
			<div><label for="description">О себе</label></div>
			<div>
				<textarea
					id="description"
					bind:value={about}
					cols="113"
					rows="10"
					placeholder="Я студент 2 курса мисис, начинающий фронтендер ищу команду для хакатона..."
				></textarea>
			</div>
		</div>
		<div>
			<h2>Выберите фотографию кота для анкеты:</h2>

			<div class="image-container">
				{#each images as image}
					<img
						src={image.src}
						alt={image.alt}
						on:click={() => handleSelection(image.src)}
						class:selected={selectedImage === image.src}
					/>
				{/each}
			</div>
		</div>

		<div class="submit">
			<input type="submit" value="Создать анкету" id="submit" on:click={create_form} />
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
			align-items: center;
			background-color: #1e1e1e;
			width: 977px;
			height: 734px;
			margin-bottom: 30px;

			.question {
				margin: 10px;
				margin-top: 15px;

				input,
				select,
				textarea {
					background-color: #1e1e1e;
					width: 929px;
					border-radius: 8px;
					color: aliceblue;
					height: 40px;
					margin-top: 10px;
				}
				.idk {
					margin-left: 9px;
				}

				#description,
				#achievements {
					height: 100px;
				}
			}
			.image-container {
				display: flex;
				justify-content: space-around;
				margin: 20px 0;
				img {
					cursor: pointer;
					width: 150px;
					height: auto;
					border: 2px solid transparent;
				}
				img.selected {
					border: 2px solid rgb(228, 114, 86);
				}
			}
			#submit {
				background-color: rgba(245, 245, 245, 1);
				height: 190%;
				width: 310px;
				border-radius: 8px;
			}
		}
	}
</style>
