import React from 'react'
import './App.css'

import NavBar from './components/NavBar'
import HeroSection from './components/HeroSection'
import ImageParent from './components/ImageParent'

const App: React.FC = () => {
  return (
    <div className="App">
      <NavBar />
      <HeroSection 
        title="Welcome to Dogr"
        subtitle="Upload a dog"
        gradient="linear-gradient(90deg, #00d2ff 0%, #3a7bd5 100%)"
        height={500} />
      <ImageParent />
    </div>
  )
}

export default App
