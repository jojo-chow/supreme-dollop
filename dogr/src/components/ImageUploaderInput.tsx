import React, { useState } from "react";

interface Props {
  apiUrl: string;
  handleImageUpload: (url: string) => void;
} 

interface ImageUploaderProps {
  onFileChange: (imageUrl: string) => void;
  onUploadSuccess: (imageUrl: string) => void;
  onUploadError: (error: string) => void;
}

const ImageUploaderInput: React.FC<ImageUploaderProps> = ({ onFileChange, onUploadSuccess, onUploadError }) => {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [imageUrl, setImageUrl] = useState<string>("");
  const [dragging, setDragging] = useState<boolean>(false);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files.length > 0) {
      const file = event.target.files[0];
      updateSelectedFile(file);
    }
  };

  const handleDrop = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    event.stopPropagation();
    setDragging(false);
    if (event.dataTransfer.files && event.dataTransfer.files.length > 0) {
      const file = event.dataTransfer.files[0];
      updateSelectedFile(file);
    }
  };

  // prevent browser from opening the file
  const handleDragOver = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    event.stopPropagation();
    setDragging(true);
  };

  // restore original functionality for browser when we leave the div
  const handleDragLeave = (event: React.DragEvent<HTMLDivElement>) => {
    event.preventDefault();
    event.stopPropagation();
    setDragging(false);
  };

  const updateSelectedFile = (file: File) => {
    setSelectedFile(file);
    setImageUrl(URL.createObjectURL(file));
    onFileChange(imageUrl);
  }
  
  return (
    <div className="max-w-xl mt-2" onDrop={handleDrop} onDragOver={handleDragOver} onDragLeave={handleDragLeave}>
      <label className="flex justify-center w-full h-32 px-4 transition 
                        bg-white border-2 border-gray-300 border-dashed 
                        rounded-md appearance-none cursor-pointer 
                        hover:border-gray-400 focus:outline-none">
        <span className="flex items-center space-x-2">
          <svg xmlns="http://www.w3.org/2000/svg" className="w-6 h-6 text-gray-600" 
                fill="none" viewBox="0 0 24 24"
                stroke="currentColor" strokeWidth="2">
            <path strokeLinecap="round" strokeLinejoin="round"
                  d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
          </svg>
          <span className="font-medium text-gray-600">
            Drop file to Attach, or <span className="text-blue-600 underline">browse</span>
          </span>
        </span>
        <input className="hidden" name="fileUpload" type="file" accept="image/*" onChange={handleFileChange}/>
      </label>
    </div>
  )
}

export default ImageUploaderInput;
