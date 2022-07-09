<script>
    export let user;
</script>

<div>
    <div class="bg-white shadow overflow-hidden sm:rounded-lg">
        <div class="px-4 py-5 sm:px-6">
            {#if user}
                    <div class="bg-white shadow overflow-hidden sm:rounded-lg">
                        <div class="px-4 py-5 sm:px-6">
                            <h3 class="text-lg leading-6 font-medium text-gray-900">
                                Github User</h3>
                            <p class="mt-1 max-w-2xl text-sm text-gray-500">
                                Personal details and info taken from Github's
                                API.</p>
                        </div>
                        <div class="border-t border-gray-200 px-4 py-5 sm:px-6">
                            <dl class="grid grid-cols-1 gap-x-4 gap-y-8 sm:grid-cols-2">
                                <div class="sm:col-span-1">
                                    <dt class="text-sm font-medium text-gray-500">
                                        Full name
                                    </dt>
                                    <dd class="mt-1 text-sm text-gray-900">{user.name}</dd>
                                </div>
                                <div class="sm:col-span-1">
                                    <dt class="text-sm font-medium text-gray-500">
                                        Company
                                    </dt>
                                    <dd class="mt-1 text-sm text-gray-900">{user.company}</dd>
                                </div>
                                <div class="sm:col-span-1">
                                    <dt class="text-sm font-medium text-gray-500">
                                        Email address
                                    </dt>
                                    <dd class="mt-1 text-sm text-gray-900">{user.email === user.email && "None supplied"}</dd>
                                </div>
                                <div class="sm:col-span-1">
                                    <dt class="text-sm font-medium text-gray-500">
                                        Location
                                    </dt>
                                    <dd class="mt-1 text-sm text-gray-900">{user.location}</dd>
                                </div>
                                <div class="sm:col-span-1">
                                    <dt class="text-sm font-medium text-gray-500">
                                        Followers
                                    </dt>
                                    <dd class="mt-1 text-sm text-gray-900">
                                        <a href="{user.followers_url}">
                                            {user.followers}
                                        </a>
                                    </dd>
                                </div>
                                <div class="sm:col-span-1">
                                    <dt class="text-sm font-medium text-gray-500">
                                        Following
                                    </dt>
                                    <dd class="mt-1 text-sm text-gray-900">{user.following}</dd>
                                </div>
                                <div class="sm:col-span-2">
                                    <dt class="text-sm font-medium text-gray-500">
                                        About
                                    </dt>
                                    <dd class="mt-1 text-sm text-gray-900">{user.bio}</dd>
                                </div>
                                <div class="sm:col-span-2">
                                    <dt class="text-sm font-medium text-gray-500">
                                        Links
                                    </dt>
                                    <dd class="mt-1 text-sm text-gray-900">
                                        <ul
                                            class="border-gray-200 rounded-md divide-y divide-gray-200">
                                            <li class="pl-3 pr-4 py-3 flex items-center justify-between text-sm">
                                                <a class="font-medium text-indigo-600 hover:text-indigo-500"
                                                   href="{user.html_url}">Github</a>
                                            </li>
                                            {#if user.twitter_username}
                                                <li class="pl-3 pr-4 py-3 flex items-center justify-between text-sm">
                                                    <a class="font-medium text-indigo-600 hover:text-indigo-500"
                                                       href="https://twitter.com/{user.twitter_username}">Twitter</a>
                                                </li>
                                            {/if}
                                            {#if user.blog}
                                                <li class="pl-3 pr-4 py-3 flex items-center justify-between text-sm">
                                                    <a class="font-medium text-indigo-600 hover:text-indigo-500"
                                                       href="https://{user.blog}">Blog</a>
                                                </li>
                                            {/if}
                                        </ul>
                                    </dd>
                                </div>
                            </dl>
                        </div>
                    </div>
            {/if}
        </div>
    </div>
</div>
