<script>
	import '../app.postcss';
	import Nav from '$lib/components/nav.svelte';
	import { setInitialClassState } from '@skeletonlabs/skeleton';

	setInitialClassState();
</script>

<Nav />
<slot />
