// src/routes/todos/[id]/+page.js

/** @type {import('@sveltejs/kit').Load} */
export async function load({ fetch, params }) {
    try {
        const id = params.id;
        const res = await fetch(`http://localhost:3000/v1/todos/${id}`);

        if (!res.ok) {
            throw new Error('failed to fetch data')
        }

        const task = await res.json();

        return {
            task
        };

    } catch (error) {
        console.error(error);
        return { task: null }
    }
}
