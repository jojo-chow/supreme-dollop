import React, { useState } from "react";
import ImageUploader from "./ImageUploader";
import ImagePreview from "./ImagePreview";

import config from "../config.json";

const ImageParent = () => {
  const apiUrl = config.apiUrl;
  const [imageUrl, setImageUrl] = useState<string | null>(null);

  const handleFileChange = (imageUrl: string) => {
    setImageUrl(imageUrl);
  }

  const handleUploadSuccess = (imageUrl: string) => {
    // TODO: do something
  };

  const handleUploadError = (error: string) => {
    // TODO: do something better than logging in the future
    console.log(error);
  };

  const handleImageUpload = (url: string) => {

  };

  return (
    <div className="flex flex-col items-center">
      <ImageUploader onFileChange={handleFileChange} onUploadSuccess={handleUploadSuccess} onUploadError={handleUploadError} />
      <ImagePreview imageUrl={imageUrl} />
    </div>
  );
};

export default ImageParent;
