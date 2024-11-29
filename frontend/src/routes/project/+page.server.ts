import { posts } from './data.js';

export function load() {
	return {
		summaries: posts.map((post) => ({
			id: post.id,
			name: post.name,
            image: post.image,
            description: post.description
		}))
	};
}
