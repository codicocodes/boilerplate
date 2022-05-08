import { useState, useEffect } from 'react'
import logo from './logo.svg'
import './App.css'

function App() {
  const [time, setTime] = useState(null)
  const [url] = useState("http://localhost:8000/v1/notifier")
  const [conn, setSSEConnection] = useState(null)
  useEffect(() => {
    console.log('connecting...')
    const sse = new EventSource(url)
    sse.onmessage = (msg) => {
      setTime(msg.data)
    };
    setSSEConnection(sse)
  }, [url])
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>Hello Vite + React!</p>
        <p>
          <span>
            last event received at: {time}
          </span>
        </p>
        <p>
          Edit <code>App.tsx</code> and save to test HMR updates.
        </p>
        <p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
          {' | '}
          <a
            className="App-link"
            href="https://vitejs.dev/guide/features.html"
            target="_blank"
            rel="noopener noreferrer"
          >
            Vite Docs
          </a>
        </p>
      </header>
    </div>
  )
}

export default App
