import { useState } from 'react'
import './App.css'
import NavBar from './components/NavBar'
import MainDog from './components/MainDog'

function App() {
  return (
    <div className="App">
      <NavBar />
      <div className='container mx-auto mt-6'>
        <MainDog />
      </div>
    </div>
  )
}

export default App
