import React, { useState } from 'react';

interface ImageUploaderProps {
  onUploadSuccess: (imageUrl: string) => void;
  onUploadError: (error: string) => void;
}

const ImageUploader: React.FC<ImageUploaderProps> = ({ onUploadSuccess, onUploadError }) => {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [imageUrl, setImageUrl] = useState<string | null>(null);
  const [dragging, setDragging] = useState<boolean>(false);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files.length > 0) {
      const file = event.target.files[0];
      setSelectedFile(file);
      const imageUrl = URL.createObjectURL(file);
      setImageUrl(imageUrl);
    }
  };

  const handleDrop = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    event.stopPropagation();
    setDragging(false);
    if (event.dataTransfer.files && event.dataTransfer.files.length > 0) {
      const file = event.dataTransfer.files[0];
      setSelectedFile(file);
      const imageUrl = URL.createObjectURL(file);
      setImageUrl(imageUrl);
    }
  };

  const handleDragOver = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    event.stopPropagation();
    setDragging(true);
  };

  const handleDragLeave = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    event.stopPropagation();
    setDragging(false);
  }

  const handleUploadClick = () => {
    const fileInput = document.getElementById('fileInput') as HTMLInputElement;
    fileInput.click();
  };

  return (
    <div
      className={`w-full max-w-xs ${
        dragging ? 'border-dashed border-2 border-blue-500' : ''
      }`}
      onDrop={handleDrop}
      onDragOver={handleDragOver}
      onDragLeave={handleDragLeave}
    >
      <label className="block text-gray-700 text-sm font-bold mb-2">
        Upload an image
      </label>
      <div className="flex items-center justify-center bg-grey-lighter">
        <label
          className="w-full flex flex-col items-center px-4 py-6 bg-white text-blue rounded-lg tracking-wide uppercase border border-blue cursor-pointer hover:bg-blue hover:text-white"
          htmlFor="fileInput"
        >
          {imageUrl ? (
            <img src={imageUrl} alt="Selected file" />
          ) : (
            <span>Choose a file or drag it here</span>
          )}
          <input
            className="hidden"
            id="fileInput"
            name="fileInput"
            type="file"
            accept="image/*"
            onChange={handleFileChange}
          />
        </label>
      </div>
      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        onClick={handleUploadClick}
      >
        Upload
      </button>
    </div>
  );
};

export default ImageUploader;