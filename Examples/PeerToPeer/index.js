import {Peer} from "https://esm.sh/peerjs@1.5.4?bundle-deps"

const friends = [
    "940168dc-9a71-4e4e-afc2-dd48af216727",
    "b8509175-4272-4d5f-a47b-34ba45da4a8b",
    "c7dbe174-02a0-425b-abd7-e390950a4f42",
    "190a649a-57fd-4ede-a743-20ca6fe56c60",
    "759bb8d4-6a32-494a-80e9-1f089e8cce25",
    "dfcd45dc-abd4-4afc-a001-aec240c1e2d7",
    "28e2c02a-09a8-480a-8e62-9a55a0971b95",
    "b4778ff4-b8ba-4a70-8888-d0c749049cc4",
    "dece3c0b-cecc-4ed2-ac20-017f33c05516",
    "00808d3e-383f-4326-b6dd-311d3bde824d",
]

// set my peer id
let peerId = localStorage.getItem('peerId');
if (peerId === null) {
    peerId = crypto.randomUUID();
    localStorage.setItem("peerId", peerId);
}
document.getElementById("peer-id").textContent = peerId;

const peer = new Peer(peerId)

peer.on('connection', function (incomingConnection) {
    incomingConnection.on('data', function (data) {
        // Will print 'hi!'
        console.log(data);
    });
});

const outgoingConnection = peer.connect('foo');
outgoingConnection.on('open', function () {
    outgoingConnection.send('hi!');
});
