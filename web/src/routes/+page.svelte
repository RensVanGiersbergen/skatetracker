<script>
	import { Geolocation } from '@capacitor/geolocation';
	import { Motion } from '@capacitor/motion';

	let loc = null;
	async function getCurrentPosition() {
		const res = await Geolocation.getCurrentPosition();
		loc = res;
	}

	let accelHandler;
	let motion = null;
	let shakiness = null;

	async function getCurrentMotion() {
		accelHandler = await Motion.addListener('accel', (event) => {
			console.log('Device motion event: ', event);
			motion = event.acceleration;
			shakiness = Math.sqrt(
				event.acceleration.x * event.acceleration.x +
					event.acceleration.y * event.acceleration.y +
					event.acceleration.z * event.acceleration.z
			);
		});
	}
</script>

<div class="container h-full mx-auto flex justify-center items-center">
	<div>
		<h1>Geolocation</h1>
		{#if loc != null}
			<p>Your location is:</p>
			<p>Latitude: {loc?.coords.latitude}</p>
			<p>Longitude: {loc?.coords.longitude}</p>
		{/if}

		<button type="button" class="btn variant-filled-primary" on:click={getCurrentPosition}>
			Get Location
		</button>

		<h1>Motion</h1>
		{#if motion != null}
			<p>Your motion is:</p>
			<p>X: {motion?.x}</p>
			<p>Y: {motion?.y}</p>
			<p>Z: {motion?.z}</p>
			<p>Shakiness: {shakiness}</p>
		{/if}

		<button type="button" class="btn variant-filled-primary" on:click={getCurrentMotion}>
			Get Motion
		</button>
	</div>
</div>
