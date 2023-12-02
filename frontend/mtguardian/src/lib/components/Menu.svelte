<script lang="ts">
  import { page } from "$app/stores";

  let active = $page.url.pathname.replace("/", "");
</script>

<nav class="bg-white border-gray-200 dark:bg-gray-700">
  <div
    class="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4"
  >
    <a
      href="https://flowbite.com/"
      class="flex items-center space-x-3 rtl:space-x-reverse"
    >
      <img
        src="https://flowbite.com/docs/images/logo.svg"
        class="h-8"
        alt="Flowbite Logo"
      />
      <span
        class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white"
        >Flowbite</span
      >
    </a>
    <button
      data-collapse-toggle="navbar-default"
      type="button"
      class="inline-flex items-center p-2 w-10 h-10 justify-center text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
      aria-controls="navbar-default"
      aria-expanded="false"
    >
      <span class="sr-only">Open main menu</span>
      <svg
        class="w-5 h-5"
        aria-hidden="true"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 17 14"
      >
        <path
          stroke="currentColor"
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M1 1h15M1 7h15M1 13h15"
        />
      </svg>
    </button>
    <div class="hidden w-full md:block md:w-auto" id="navbar-default">
      <ul
        class="font-large flex flex-col p-4 md:p-0 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 rtl:space-x-reverse md:mt-0 md:border-0 md:bg-white dark:bg-gray-700 md:dark:bg-gray-700 dark:border-gray-600"
      >
        <li>
          <a
            href="/"
            on:click={() => (active = "")}
            class:active={active === ""}
            class="block py-2 px-3 text-white rounded md:bg-transparent md:p-0 dark:text-white"
            aria-current="page">Home</a
          >
        </li>
        <li>
          <a
            href="/prescriptions"
            on:click={() => (active = "prescriptions")}
            class:active={active === "prescriptions"}
            class="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent"
            >Medication</a
          >
        </li>

        <li>
          <a
            href="/events"
            on:click={() => (active = "events")}
            class:active={active === "events"}
            class="block py-2 px-3 text-gray-900 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent"
            >Events</a
          >
        </li>
      </ul>
    </div>
  </div>
</nav>

<style>
  .active {
    color: blue;
  }
</style>
