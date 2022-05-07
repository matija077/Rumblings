
const sendRequestButton = document.getElementById('send-request');
sendRequestButton.addEventListener('click', sendRequestHandler)

function sendRequestHandler() {
    for (let i = 0; i < 10; i++) {
        sendRequest()
    }
}

function sendRequest() {
    fetch('http://localhost:8090/get') .then(response => response.json())
    .then(data => console.log(data))
    .catch(err => console.error(err))
    .finally(() => console.warn('request completed'))
}

console.dir(sendRequestButton)
console.log('ha ha')
