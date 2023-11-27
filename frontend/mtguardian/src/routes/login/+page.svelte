<script lang="ts">
  let register = false;
</script>

<div class="background">
  <div class="flex flex-col justify-center items-center h-screen">
    {#if !register}
      <div class="text-5xl pb-3 text-white fade-in-login">Login</div>
      <form
        method="POST"
        class="bg-gray-700 p-10 rounded-xl fade-in-form slide-up"
      >
        <label class="block text-lg font-medium text-white" for="email">
          Email
        </label>
        <input class="input" id="email" name="email" />

        <label class="block text-lg font-medium text-white" for="password">
          Password
        </label>
        <input class="input" id="password" type="password" name="password" />
        <div>
          <button class="btn btn-primary btn-wide mt-6" type="submit"
            >{register ? "Register" : "Login"}</button
          >
        </div>
      </form>
    {:else}
      <div class="text-5xl pb-3 text-white fade-in-login">Register</div>
      <form
        method="POST"
        class="bg-gray-700 p-10 rounded-xl fade-in-form slide-up"
      >
        <label class="block text-lg font-medium text-white" for="email">
          Email
        </label>
        <input class="input" id="email" name="email" />

        <label class="block text-lg font-medium text-white" for="password">
          Password
        </label>
        <input class="input" id="password" type="password" name="password" />

        <label class="block text-lg font-medium text-white" for="password">
          Confirm Password
        </label>
        <input class="input" id="password" type="password" name="password" />
        <div>
          <button class="btn btn-primary btn-wide mt-6" type="submit"
            >{register ? "Register" : "Login"}</button
          >
        </div>
      </form>
    {/if}
    {#if !register}
      <div class="mt-2">
        <button on:click={() => (register = true)}>Register</button>
      </div>
    {:else}
      <div class="mt-2">
        <button on:click={() => (register = false)}>Login</button>
      </div>
    {/if}
  </div>
</div>

<style>
  .background {
    background: linear-gradient(to right, #1a1a1a, #2d2d2d, #000000);
  }

  .fade-in-login {
    opacity: 0;
    animation: fadeInAnimation 2000ms forwards;
  }

  .fade-in-form {
    opacity: 0;
    animation: fadeInAnimation 4000ms forwards;
  }

  @keyframes fadeInAnimation {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }

  .slide-up {
    opacity: 0;
    transform: translateY(3vh);
    animation: slideUpAnimation 1000ms forwards;
  }

  @keyframes slideUpAnimation {
    from {
      opacity: 0;
      transform: translateY(20px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
</style>
