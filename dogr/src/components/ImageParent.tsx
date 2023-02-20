import { useState } from 'react';
import ImageUploader from './ImageUploader';
import ImageDisplay from './ImageDisplay';

const ImageParent = () => {
  const [imageUrl, setImageUrl] = useState<string | null>(null);

  const handleUploadSuccess = (imageUrl: string) => {
    setImageUrl(imageUrl);
  };

  const handleUploadError = (error: string) => {
    // TODO: do something better than logging in the future
    console.log(error);
  };

  return (
    <div className="flex flex-row">
      <ImageUploader onUploadSuccess={handleUploadSuccess} onUploadError={handleUploadError} />
      <ImageDisplay imageUrl={imageUrl} />
    </div>
  )
};

export default ImageParent;