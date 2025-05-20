<script>
	import { onMount } from "svelte";
	import { page } from "$app/stores";

	/**
	 * @typedef {Object} Todo
	 * @property {number} id
	 * @property {string} title
	 * @property {string} description
	 * @property {boolean} completed
	 * @property {string} created_at
	 * @property {string} due_date
	 * @property {'low' | 'medium' | 'high'} priority
	 * @property {string[]} tags
	 * @property {string} updated_at
	 * @property {number} version
	 */

	/** @type {Todo | null}  */
	let todo;

	onMount(async () => {
		/** {Response} */
		const res = await fetch(`http://localhost:3000/v1/todos/69`);
		/** {{task: Todo}} */
		const data = await res.json();

		todo = data.task;
	});
</script>

<h1 class="text-3xl font-bold text-center">
	Svelte + Go + PostgreSQL Todo App
</h1>
<p class="pt-2 text-xs text-center">
	This application is built using the SvelteKit framework and Golang using the
	PostgreSQL database
</p>

{#if todo}
	<div class="p-4">
		<h1 class="text-2xl font-bold">{todo.title}</h1>
		<p class="text-sm text-gray-500">Priority: {todo.priority}</p>
		<p class="mt-2">{todo.description}</p>
		<p class="mt-2 text-sm">
			Due: {new Date(todo.due_date).toLocaleString()}
		</p>
		<p class="mt-2 text-sm">Tags: {todo.tags.join(", ")}</p>
	</div>
{:else}
	<p class="p-4">Loading...</p>
{/if}
