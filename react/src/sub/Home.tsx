import { useState } from 'react'
import { Link } from 'react-router-dom'
import '../App.css'

function Home() {
  const [count, setCount] = useState(0)

  return (
    <>
      <div>
        <Link to="/home">Home</Link><br />
        <Link to="/page1">Page1</Link><br />
        <Link to="/page2">Page2</Link><br />
        <Link to="/clip/list">ClipList</Link><br />
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>

    </>
  )
}

export default Home