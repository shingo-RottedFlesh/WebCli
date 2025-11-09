import './App.css'
import {BrowserRouter, Routes, Route, Link} from 'react-router-dom'
import Page1 from './sub/Page1.tsx'
import Page2 from './sub/Page2.tsx'
import Home from './sub/Home.tsx'
import Login from './login/LoginForm.tsx'

function App() {
  return (
    <>
      <BrowserRouter>
        {/* <Link to="/">App</Link><br/> */}
        <Link to="/home">Home</Link><br/>
        <Link to="/page1">Page1</Link><br/>
        <Link to="/page2">Page2</Link><br/>
        <Link to="/login">login</Link><br/>
        <Routes>
          {/* <Route path="/" element={<App />} /> */}
          <Route path="/home" element={<Home />} />
          <Route path="/page1" element={<Page1 />} />
          <Route path="/page2" element={<Page2 />} />
          <Route path="/login" element={<Login />} />
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
