import React from "react";
import "./App.css";

import NavBar from "./components/NavBar";
import HeroSection from "./components/HeroSection";
import ImageParent from "./components/ImageParent";
import Footer from "./components/Footer";

const App: React.FC = () => {
  return (
    <div className="App">
      <NavBar />
      <div className="flex-grow">
        <HeroSection 
          title="Welcome to Dogr"
          subtitle="Upload a dog"
          gradient="linear-gradient(90deg, #00d2ff 0%, #3a7bd5 100%)"
          height={500} />
        <ImageParent />
      </div>
      <Footer />
    </div>
  )
}

export default App
