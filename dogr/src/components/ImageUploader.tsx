import React, { useState } from "react";
import axios from "axios";
import { useDropzone } from "react-dropzone";
import { AiOutlineCloudUpload } from "react-icons/ai";

interface Props {
  apiUrl: string;
  handleImageUpload: (url: string) => void;
}

const ImageUploader: React.FC<Props> = ({ apiUrl, handleImageUpload }) => {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [isUploading, setIsUploading] = useState<boolean>(false);

  const handleFileSelection = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files && e.target.files[0];
    setSelectedFile(file);
  }

  const handleFileDrop = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    const file = e.dataTransfer.files && e.dataTransfer.files[0];
    setSelectedFile(file);
  }

  const handleFileUpload = async () => {
    if (!selectedFile) return;

    setIsUploading(true);

    try {
      const formData = new FormData();
      formData.append("image", selectedFile);

      const response = await axios.post(apiUrl, formData);

      if (response && response.data && response.data.url) {
        handleImageUpload(response.data.url);
      }
    } catch (error) {
      console.error(error);
    } finally {
      setIsUploading(false);
    }
  }

  return (
    <div className="relative">
      <label htmlFor="image-upload-input">
        <div
          className="border-dashed border-2 border-gray-400 h-64 flex justify-center items-center cursor-pointer"
          onDragOver={(e) => e.preventDefault()}
          onDrop={handleFileDrop}
        >
          {selectedFile ? (
            <img
              className="h-full w-full object-contain"
              src={URL.createObjectURL(selectedFile)}
              alt="Preview"
            />
          ) : (
            <span>Drag and drop or click to upload</span>
          )}
        </div>
      </label>

      <input
        id="image-upload-input"
        type="file"
        className="hidden"
        onChange={handleFileSelection}
      />

      <button
        className="absolute bottom-0 right-0 bg-blue-500 text-white py-2 px-4 rounded-md m-4 disabled:opacity-50"
        onClick={handleFileUpload}
        disabled={!selectedFile || isUploading}
      >
        {isUploading ? "Uploading..." : "Upload"}
      </button>
    </div>
  );

};

export default ImageUploader;
