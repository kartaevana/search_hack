import { api } from "../../api.js";

export async function load({ params }: any) {
	//id команды!!
	const id = params.id;

	const response = await fetch(api + `/team/{id}?id=${id}`, {
		headers: {
			"Content-Type": "application/json"
		}
	});
	const obj = await response.json();
	const data = obj.team; // Извлекаем команду из ответа
	console.log(data);

	return {
		team: data
	};
}
