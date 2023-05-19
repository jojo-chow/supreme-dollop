import { useState } from "react";
import { render } from "react-dom";

interface ImagePreviewProps {
  imageUrl: string | null;
}

const ImagePreview: React.FC<ImagePreviewProps> = ({ imageUrl }) => {
  
  const render = () => {
    if (imageUrl) {
      return <img src={imageUrl} alt="Uploaded file" />;
    }
    return <div>No image uploaded.</div>
  }

  return render();
};

export default ImagePreview;
