document.addEventListener('DOMContentLoaded', function () {
  const stationsDataElement = document.getElementById('station-data')
  const stationsData = stationsDataElement.getAttribute('data-stations')
  const lineStringsData = stationsDataElement.getAttribute('data-linestrings')

  const stations = JSON.parse(stationsData)
  let lineStrings = JSON.parse(lineStringsData)

  lineStrings = lineStrings.map((line) =>
    JSON.parse(line)
      .flat()
      .map(([lon, lat]) => [lat, lon])
  )

  const initialMapPosition = stations[0]
    ? [stations[0].lat, stations[0].lon]
    : [51.5074, -0.1278]

  const map = L.map('map').setView(initialMapPosition, 12)

  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution:
      '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
  }).addTo(map)

  const polyline = L.polyline(lineStrings, { color: 'red' }).addTo(map)
  map.fitBounds(polyline.getBounds())

  stations.forEach((station) => {
    const marker = L.marker([station.lat, station.lon])
      .addTo(map)
      .bindPopup(station.name)

    marker.on('click', function () {
      fetch(`/stop/arrivals/${station.stationId}`)
        .then((res) => {
          return res.json()
        })
        .then((arrivals) => {
          displayArrivals(station.name, arrivals)
        })
        .catch((error) => {
          console.error('error fetching arrivals:', error)
        })
    })
  })
})

function displayArrivals(stationName, arrivals) {
  const arrivalsContainer = document.getElementById('arrivals__container')
  arrivalsContainer.innerHTML = ''

  const stationTitle = document.createElement('h2')
  stationTitle.textContent = `Arrivals at ${stationName}`
  arrivalsContainer.appendChild(stationTitle)

  arrivals.forEach((arrival) => {
    const arrivalCard = document.createElement('div')
    arrivalCard.classList.add('arrival__card')

    const lineName = document.createElement('h3')
    lineName.textContent = `${arrival.lineName} line`
    arrivalCard.appendChild(lineName)

    const destinationName = document.createElement('p')
    destinationName.classList.add('destination')
    destinationName.textContent = `Destination: ${arrival.destinationName}`
    arrivalCard.appendChild(destinationName)

    const platformName = document.createElement('p')
    platformName.classList.add('platform')
    platformName.textContent = `${
      !arrival.platformName.toLowerCase().includes('platform')
        ? `Platform ${arrival.platformName}`
        : `${arrival.platformName}`
    }`
    arrivalCard.appendChild(platformName)

    const expectedArrival = document.createElement('p')
    expectedArrival.classList.add('expected-arrival')
    expectedArrival.textContent = `${arrival.expectedArrival}`
    arrivalCard.appendChild(expectedArrival)

    const currentLocation = document.createElement('p')
    currentLocation.classList.add('current-location')
    currentLocation.textContent = `Current Location: ${arrival.currentLocation}`
    arrivalCard.appendChild(currentLocation)

    const towards = document.createElement('p')
    towards.textContent = `Heading Towards: ${arrival.towards}`
    arrivalCard.appendChild(towards)

    arrivalsContainer.append(arrivalCard)
  })
}
