import './App.css'
import { Routes, Route } from 'react-router-dom'
import Page1 from './sub/Page1.tsx'
import Page2 from './sub/Page2.tsx'
import Home from './sub/Home.tsx'
import Login from './features/login.tsx'
import ClipList from './features/ClipList.tsx'
import RequireAuth from './components/RequireAuth.tsx'
import React from 'react';

const App: React.FC = () => {
  return (
    <div>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route
          path="/home"
          element={
            <RequireAuth>
              <Home />
            </RequireAuth>
          }
        />
        <Route
          path="/page1"
          element={
            <RequireAuth>
              <Page1 />
            </RequireAuth>
          }
        />
        <Route
          path="/page2"
          element={
            <RequireAuth>
              <Page2 />
            </RequireAuth>
          }
        />
        <Route
          path="/clip/list"
          element={
            <RequireAuth>
              <ClipList />
            </RequireAuth>
          }
        />
      </Routes>
    </div>
  );
}



export default App
