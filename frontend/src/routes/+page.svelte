<script lang="ts">
	import { goto } from "$app/navigation";
	import "../app.css";
	function go_sign_up() {
		goto("/sign_up");
	}
	function go_sign_in() {
		goto("/sign_in");
	}

//  смена цвета для кнопки регистрации
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

	//  смена цвета для кнопки входа
	let color2 = "#1e1e1e";
	let font_color2 = "#ffffff";

	function handleMouseOver2(e) {
		color2 = "#ffffff";
		font_color2 = "black"
	}
	function handleMouseOut2(e) {
		color2 = "#1e1e1e";
		font_color2 = "#ffffff";
	}




</script>

<header>
	<a href="/">
		<img height="24px" src="/cover.png" alt="" style="margin-left:15px" />
	</a>
</header>

<main>
	<div class="logo">
		<h1>Dream team</h1>
		<h2>Присоединяйтесь к нам!</h2>
		<div>
			<button
				on:click={go_sign_up}
				id="sign_up"
				style="background-color: {color1}; color: {font_color1}"
				on:mouseover={handleMouseOver}
				on:mouseout={handleMouseOut}
			>
				Регистрация</button
			>
			<button on:click={go_sign_in} id="sign_in" style="background-color: {color2}; color: {font_color2}"
			on:mouseover={handleMouseOver2}
			on:mouseout={handleMouseOut2} > Вход</button>
		</div>
	</div>
</main>

<style lang="scss">
	main {
		display: flex;
		justify-content: center;
		.logo {
			margin-top: 50px;
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

			button {
				margin-top: 20px;
				height: 32px;
				padding: 3px;
				border-radius: 8px;
			}
			#sign_up {
				// background-color: rgba(44, 44, 44, 1);
				// color: #ffffff;
				width: 118px;
			}
			#sign_in {
				width: 55px;
			}
		}
	}
</style>
