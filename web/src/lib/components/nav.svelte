<script>
	import { LightSwitch, TabGroup, TabAnchor } from '@skeletonlabs/skeleton';
	import { page } from '$app/stores';
	let tabs = [
		{ name: 'Home', href: '/' },
		{ name: 'Tracking', href: '/tracking' }
	];
</script>

<div class="bg-surface-100-800-token flex items-center p-4 space-x-2">
	<TabGroup
		active="variant-ghost-primary"
		hover="hover:variant-soft-primary"
		flex="flex-1 lg:flex-none"
		rounded=""
		border=""
		class="w-full"
	>
		{#each tabs as tab}
			<TabAnchor class="mx-1" href={tab.href} selected={$page.url.pathname === tab.href}
				>{tab.name}</TabAnchor
			>
		{/each}
	</TabGroup>
	<div>
		<LightSwitch />
	</div>
</div>
