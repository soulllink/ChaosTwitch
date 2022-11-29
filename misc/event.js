const url = 'ws://localhost:8080/echo'
const connection = new WebSocket(url)


connection.onerror = error => {
  console.log(`WebSocket error: ${error}`)
}

// connection.onopen = () => {
//   connection.send('Web client is refreshed')
// }

connection.onmessage = e => {
  
    let aurl = './media/' + e.data + '.mp3';
    console.log(e.data);
    let audio = new Audio(aurl);
    audio.load();
    audio.play();
  
}