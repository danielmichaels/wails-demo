<script>
    import {GetGHUser} from "../../wailsjs/go/main/App.js";
    import UserCard from "./UserCard.svelte";

    let ghUsername = "rs"
    let response
    let err


    function getGHUser() {
        GetGHUser(ghUsername)
            .then(result => response = result)
            .catch(e => err = e)
    }
</script>

<div>
    <div class="min-h-full flex flex-col justify-center py-12 sm:px-6 lg:px-8">
        <div class="sm:mx-auto sm:w-full sm:max-w-md">
            <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
               Search for a GitHub user</h2>
            <p class="mt-2 text-center text-sm text-gray-600">
                Example users: <i>rs</i>, <i>rsc</i> and <i>alexellis</i>
            </p>
        </div>
        
        <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
            <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
                <div class="space-y-6" >
                    <div>
                        <label for="text" class="block text-sm font-medium text-gray-700">
                            Username</label>
                        <div class="mt-1">
                            <input bind:value={ghUsername}
                                    id="text" name="text" type="text"
                                   placeholder="rs"
                                   autocomplete="text" required
                                   class="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
                        </div>
                    </div>
                    <div>
                        <button on:click={getGHUser}
                                class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                            Search
                        </button>
                    </div>
                </div>
            </div>
        </div>
        {#if response}
            {response.name}
            <UserCard user={response} />
        {/if}
        {#if err}
            {err}
        {/if}
    </div>

</div>

