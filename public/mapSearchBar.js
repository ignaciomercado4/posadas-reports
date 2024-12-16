L.Control.geocoder({
    defaultMarkGeocode: false,
    placeholder: "Search address in Posadas...",
    geocoder: L.Control.Geocoder.nominatim({
        geocodingQueryParams: {
            viewbox: "-57.0,-28.0,-54.0,-26.0",
            countrycodes: 'ar',
            limit: 10,
            bounded: 0
        }
    })
}).on('markgeocode', function(e) {
    console.log('Geocode result:', e.geocode);  

    if (e.geocode && e.geocode.center) {
        map.eachLayer(function(layer) {
            if (layer instanceof L.Marker && layer !== newReportMarker) {
                map.removeLayer(layer);
            }
        });

        const latlng = e.geocode.center;
        const searchMarker = L.marker(latlng, {
            icon: L.divIcon({
                className: 'search-marker',
                html: `
                    <div class="relative w-10 h-10 bg-blue-600 rounded-full border-4 border-white shadow-lg animate-pulse">
                        <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-3 h-3 bg-white rounded-full"></div>
                    </div>
                `,
                iconSize: [40, 40],
                iconAnchor: [20, 20]
            })
        }).addTo(map);
        
        searchMarker.bindPopup(e.geocode.name).openPopup();
        map.setView(latlng, 16);
    } else {
        alert('No se encontraron resultados para esta dirección');
    }
}).addTo(map);