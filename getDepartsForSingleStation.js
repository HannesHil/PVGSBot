const createClient = require('hafas-client')
const vbbProfile = require('hafas-client/p/insa')
let departs = "a"
const client = createClient(vbbProfile, 'my-awesome-program')
client.departures('461', {duration: 3}).then(function(stop){departs=stop}).catch(console.error)
