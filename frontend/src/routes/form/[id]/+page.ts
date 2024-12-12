import { api } from "../../api.js";

export async function load({ params }: any) {
	const id = params.id;

	const response = await fetch(api + `/user/{id}?id=${id}`, {
		headers: {
			"Content-Type": "application/json"
		}
	});
	const obj = await response.json();
	const data = obj.user; // Извлекаем пользователя из ответа
    console.log(data);

	return {
			user: data
	};
}
