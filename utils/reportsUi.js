const posadas = [-27.3621, -55.9045];

const map = L.map('map').setView(posadas, 12);

L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '© OpenStreetMap contributors'
}).addTo(map);

const categoryColors = {
    'Infrastructure': 'bg-red-500',
    'Environment': 'bg-green-500',
    'Public Safety': 'bg-blue-500',
    'Transportation': 'bg-purple-500',
    'Garbage Disposal': 'bg-orange-700',
    'Road Marking/Signaling': 'bg-lime-500',
    'Vandalism': 'bg-fuchsia-500',
    'Other': 'bg-yellow-500'
};

function createMarkerIcon(category) {
    const color = categoryColors[category] || 'bg-gray-500';
    
    return L.divIcon({
        className: 'custom-div-icon',
        html: `
            <div class="relative w-8 h-8 ${color} rounded-full border-2 border-white shadow-md">
                <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-4 h-4 bg-white rounded-full"></div>
            </div>
        `,
        iconSize: [32, 32],
        iconAnchor: [16, 16]
    });
}


const reportModal = document.getElementById('reportModal');
const closeModalBtn = document.getElementById('closeModal');
const addReportBtn = document.getElementById('addReportBtn');
const addressInput = document.getElementById('address');
let newReportMarker = null;

function fetchAddress(lat, lng) {
    const url = `https://nominatim.openstreetmap.org/reverse?format=json&lat=${lat}&lon=${lng}&zoom=18&addressdetails=1`;
    
    fetch(url)
        .then(response => response.json())
        .then(data => {
            let address = '';
            if (data.address) {
                const details = data.address;
                address = [
                    details.road || '',
                    details.house_number || '',
                    details.city || details.town || details.village || '',
                    details.state || '',
                    details.postcode || ''
                ].filter(Boolean).join(', ');
            }
            
            addressInput.value = address || `${lat.toFixed(6)}, ${lng.toFixed(6)}`;
        })
        .catch(error => {
            console.error('Error fetching address:', error);
            addressInput.value = `${lat.toFixed(6)}, ${lng.toFixed(6)}`;
        });
}

addReportBtn.addEventListener('click', function() {
    const center = map.getCenter();
    
    if (newReportMarker) {
        map.removeLayer(newReportMarker);
    }

    newReportMarker = L.marker(center).addTo(map);

    fetchAddress(center.lat, center.lng);

    document.getElementById('positionX').value = center.lat.toFixed(6);
    document.getElementById('positionY').value = center.lng.toFixed(6);

    reportModal.classList.remove('hidden');
    reportModal.classList.add('flex');
});

map.on('click', function(e) {
    const lat = e.latlng.lat;
    const lng = e.latlng.lng;

    if (newReportMarker) {
        map.removeLayer(newReportMarker);
    }

    newReportMarker = L.marker([lat, lng]).addTo(map);

    fetchAddress(lat, lng);

    document.getElementById('positionX').value = lat.toFixed(6);
    document.getElementById('positionY').value = lng.toFixed(6);

    reportModal.classList.remove('hidden');
    reportModal.classList.add('flex');
});

closeModalBtn.addEventListener('click', function() {
    if (newReportMarker) {
        map.removeLayer(newReportMarker);
        newReportMarker = null;
    }

    reportModal.classList.remove('flex');
    reportModal.classList.add('hidden');
});

document.getElementById('newReportForm').addEventListener('submit', function(event) {
    event.preventDefault();
    
    const posX = document.getElementById('positionX').value;
    const posY = document.getElementById('positionY').value;

    const formData = {
        title: event.target.title.value,
        category: event.target.category.value,
        description: event.target.description.value,
        address: event.target.address.value,
        positionX: parseFloat(posX),
        positionY: parseFloat(posY)
    };

    fetch('/reports/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(formData)
    })
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
        alert('Report created successfully!');
        location.reload();
    })
    .catch((error) => {
        console.error('Error:', error);
        alert('Failed to create report');
    });
});