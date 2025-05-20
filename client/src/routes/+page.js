
/** @type {import('./$types').PageLoad} */
export async function load({ fetch, params }) {
    try {
        const res = await fetch(`http://localhost:3000/v1/todos/69`);

        if (!res.ok) {
            throw new Error('failed to fetch data')
        }

        const task = await res.json();

        return { task };

    } catch (error) {
        console.error(error);
        return { task: null }
    }
}
