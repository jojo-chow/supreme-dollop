import React, { useState } from "react";
import ImageUploader from "./ImageUploader";
import ImageDisplay from "./ImageDisplay";

import config from "../config.json";

const ImageParent = () => {
  const apiUrl = config.apiUrl;
  const [previewUrl, setPreviewUrl] = useState<string | null>(null);

  const handleImageUpload = (url: string) => {
    setPreviewUrl(url);
  }

  return (
    <div className="flex flex-col items-center space-y-4">
      <ImageDisplay previewUrl={null} />
      <ImageUploader apiUrl={apiUrl} onImageUpload={handleImageUpload}/>
    </div>
  );
};

export default ImageParent;
