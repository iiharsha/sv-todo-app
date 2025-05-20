<!--src/routes/todos/[id]/+page.svelte -->
<script>
    /**
     * @typedef {import('$lib/types').Todo} Todo
     */

    /** @type {{ data: { task: Todo }}} */
    let { data } = $props();
    $inspect(data);
    let todo = data.task.task; //fix this what is data.task.task??
</script>

<div class="p-4">
    <h1 class="text-2xl font-bold">{todo.title}</h1>
    <p class="text-sm text-orange-300">Priority: {todo.priority}</p>
    <p class="mt-2">{todo.description}</p>
    <p class="mt-2 text-sm">
        Due: {new Date(todo.due_date).toLocaleString()}
    </p>
    <p class="mt-2 text-sm">Tags: {todo.tags.join(", ")}</p>
</div>
