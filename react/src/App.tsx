import './App.css'
import {Routes, Route} from 'react-router-dom'
import Page1 from './sub/Page1.tsx'
import Page2 from './sub/Page2.tsx'
import Home from './sub/Home.tsx'
import Login from './login/LoginForm.tsx'
import React from 'react';

const App: React.FC = () => {
  return (
    <div>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/home" element={<Home />} />
        <Route path="/page1" element={<Page1 />} />
        <Route path="/page2" element={<Page2 />} />
      </Routes>
    </div>
  );
}



export default App
