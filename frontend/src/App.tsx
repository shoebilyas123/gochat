import React, { useState } from 'react';
import './App.css';
import { useEffect } from 'react';
import { connect, sendMessage } from './api';
import Header from './components/Header';
import ChatHistory from './components/ChatHistory';

function App() {
  const [chatHistory, setChatHistory] = useState<any>([]);

  useEffect(() => {
    connect((msg: any) => {
      console.log('New Message');
      setChatHistory((prev: any) => [...prev, msg]);
      console.log(chatHistory);
    });
  }, []);

  const sendMsgHandler = () => {
    sendMessage('Hello');
  };

  return (
    <div>
      <Header />
      <button onClick={sendMsgHandler}>Send Message</button>
      <ChatHistory chatHistory={chatHistory} />
    </div>
  );
}

export default App;
