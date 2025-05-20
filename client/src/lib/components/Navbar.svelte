<script>
    /**
     * @typedef {Object} Navlink
     * @property {string} name
     * @property {string} path
     */

    /** @type {Navlink[]}*/
    export const navLinks = [
        { name: "Home", path: "/" },
        { name: "AddTodo", path: "/api/add" },
        { name: "ShowTodos", path: "/api/todos" },
    ];
</script>

<nav
    class="flex gap-4 p-4 items-center justify-center w-full border-b border-b-coolgray mb-4"
>
    {#each navLinks as link}
        <a href={link.path} class="text-lightgray hover:text-orange-400">
            {link.name}
        </a>
    {/each}
</nav>
