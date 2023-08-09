let socket = new WebSocket('ws://localhost:8080/ws');

const connect = () => {
  console.log('Attempting connection');

  socket.onopen = () => {
    console.log('Successfullu connected!');
  };

  socket.onmessage = (msg) => {
    console.log('MESSAGE RECIEVED:', msg);
  };

  socket.onclose = () => {
    console.log('Socket closed connection');
  };

  socket.onerror = (err) => {
    console.log('web socket error', err);
  };
};

const sendMessage = (msg: string) => {
  console.log('sending message', msg);
  socket.send(msg);
};

export { connect, sendMessage };
