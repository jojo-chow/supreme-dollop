import React, { useState } from "react";
import ImageUploaderInput from "./ImageUploaderInput";
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
      <ImageUploaderInput onFileChange={handleFileChange} onUploadSuccess={handleUploadSuccess} onUploadError={handleUploadError} />
      <ImagePreview imageUrl={imageUrl} />
      <button className="mt-2 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full" type="button">Submit</button>
    </div>
  );
};

export default ImageParent;
