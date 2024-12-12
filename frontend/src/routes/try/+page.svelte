<script lang="ts">
	import { api } from "../api.js";
	export async function load() {
		const response = await fetch(api + `/user/{id}?id=5`, {
			// Используем id в URL
			headers: {
				"Content-Type": "application/json"
			}
		});

		if (!response.ok) {
			throw new Error(`Ошибка загрузки: ${response.status}`); // Обработка ошибки, если запрос не успешен
		}

		const obj = await response.json();
		const data = obj.user; // Извлекаем пользователя из ответа
		console.log(data);
	}
    load();
</script>

meoq
