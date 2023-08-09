import React from 'react';
import './App.css';
import { useEffect } from 'react';
import { connect, sendMessage } from './api';

function App() {
  useEffect(() => {
    connect();
  }, []);

  const sendMsgHandler = () => {
    sendMessage('Hello');
  };

  return (
    <>
      <button onClick={sendMsgHandler}>Send Message</button>
    </>
  );
}

export default App;
