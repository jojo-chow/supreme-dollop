import { useState } from "react";
import { render } from "react-dom";

interface ImagePreviewProps {
  imageUrl: string | null;
}

const ImagePreview: React.FC<ImagePreviewProps> = ({ imageUrl }) => {
  
  const render = () => {
    if (imageUrl) {
      return <img className="mt-2" src={imageUrl} alt="Uploaded file" />;
    }
    return <div className="mt-2">No image added.</div>
  }

  return render();
};

export default ImagePreview;
