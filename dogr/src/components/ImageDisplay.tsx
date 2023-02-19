import { useState } from 'react';
import { render } from 'react-dom';

interface ImageDisplayProps {
  imageUrl: string | null;
}

const ImageDisplay: React.FC<ImageDisplayProps> = ({ imageUrl }) => {
  
  const render = () => {
    if (imageUrl) {
      return <img src={imageUrl} alt="Uploaded file" />;
    }
    return <div>No image uploaded.</div>
  }

  return render();
};

export default ImageDisplay;