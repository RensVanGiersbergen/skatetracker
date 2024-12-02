<script>
	import { onMount } from 'svelte';
	import L from 'leaflet';
	import 'leaflet/dist/leaflet.css';

	import { Geolocation } from '@capacitor/geolocation';

	var loc = null;

	async function getCurrentPosition() {
		try {
			const res = await Geolocation.getCurrentPosition({ enableHighAccuracy: true });
			loc = res;
		} catch (e) {
			console.log(e);
		}
	}

	onMount(async () => {
		await getCurrentPosition();

		if (loc != null) {
			var map = L.map('map').setView([loc.coords.latitude, loc.coords.longitude], 13);
		} else {
			var map = L.map('map').setView([52.1436278895767, 5.543447534013997], 8);
		}
		L.tileLayer('https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png', {
			attribution: `&copy;<a href="https://www.openstreetmap.org/copyright" target="_blank">OpenStreetMap</a>`,
			minZoom: 10
		}).addTo(map);
	});
</script>

<div id="mapContainer">
	<div id="map"></div>
</div>

<style>
	#map {
		height: calc(100vh - 100px);
	}
</style>
